import axios from "axios";
import { Trash2 } from "lucide-react";
import React from "react";
import { useNavigate } from "react-router-dom";

import { useAuth } from "../context/UseAuth.tsx";
import { useSong } from "../context/UseSong.tsx";

export const PlaylistTable = ({
  updateSong,
  detail,
  index,
  currUser,
}: {
  updateSong: () => void;
  detail: PlaylistDetail;
  index: number;
  currUser: User;
}) => {
  const { user } = useAuth();
  const navigate = useNavigate();
  const { updatePlaylist } = useSong();

  const handleDelete = (e: React.MouseEvent<SVGSVGElement>) => {
    e.stopPropagation();
    if (user == null) return;
    axios
      .delete(
        "http://localhost:4000/auth/playlist-detail?id=" +
          detail.playlistDetailId +
          "&userId=" +
          user.user_id +
          "&detId=" +
          detail.playlistId,
        {
          withCredentials: true,
        },
      )
      .then((res) => {
        console.log(res);
        updateSong();
        updatePlaylist();
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  };

  return (
    <div
      className={"playlistTable"}
      onClick={() => {
        navigate("/track/" + detail.songId);
      }}
    >
      <div className={"title"}>
        <p>{index + 1}. </p>
        <img src={detail.song.album.banner} alt="Song Cover" />
        <div>
          <h3>{detail.song.title}</h3>
          <p>{detail.song.artist.user.username}</p>
        </div>
      </div>
      <p>{detail.song.album.title}</p>
      <p>
        {new Date(detail.dateAdded).toLocaleDateString("en-US", {
          year: "numeric",
          month: "long",
          day: "numeric",
        })}
      </p>
      <div>
        <p>
          {Math.floor(detail.song.duration / 60)}:
          {Math.floor(detail.song.duration % 60)
            .toString()
            .padStart(2, "0")}
        </p>
      </div>
      {user && user.user_id == currUser.user_id && (
        <Trash2 className={"trash"} onClick={handleDelete} />
      )}
    </div>
  );
};
