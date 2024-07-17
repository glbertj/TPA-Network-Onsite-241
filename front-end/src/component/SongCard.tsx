import { Play } from "lucide-react";
import React from "react";
import { useNavigate } from "react-router-dom";

import { useAuth } from "../context/UseAuth.tsx";
import { useSong } from "../context/UseSong.tsx";

export const SongCard = ({ songs, play }: { songs: Song; play: boolean }) => {
  const navigate = useNavigate();
  const { user } = useAuth();
  const { enqueue, waitingSong, song } = useSong();

  const handlePlayClick = (
    e: React.MouseEvent<HTMLSpanElement>,
    songs: Song,
  ) => {
    e.stopPropagation();
    if (waitingSong?.find((s) => s.songId == songs.songId)) return;
    if (song?.songId == songs.songId) return;
    enqueue(songs, user);
  };

  const handleNavigate = () => {
    navigate("/track/" + songs.songId);
  };
  return (
    <div className={"card"} onClick={handleNavigate}>
      <div className={"cardImage"}>
        <img src={songs.album.banner} alt={"placeholder"} />
        {play && (
          <span
            className={"play"}
            onClick={(e) => {
              handlePlayClick(e, songs);
            }}
          >
            <Play />
          </span>
        )}
      </div>
      <div className={"cardContent"}>
        <h3>{songs.title}</h3>
        <p>{songs.artist.user.username}</p>
      </div>
    </div>
  );
};
