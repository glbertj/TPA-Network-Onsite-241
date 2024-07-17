import type { AxiosResponse } from "axios";
import axios from "axios";
import { Clock, Play, Trash2 } from "lucide-react";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

import { ControlMusic } from "../component/ControlMusic.tsx";
import { Footer } from "../component/Footer.tsx";
import { Main } from "../component/Main.tsx";
import { PlaylistTable } from "../component/PlaylistTable.tsx";
import { RichText } from "../component/RichText.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { useSong } from "../context/UseSong.tsx";

export const PlaylistPage = () => {
  const { id } = useParams();
  const { user, authenticated } = useAuth();
  const { updatePlaylist } = useSong();
  const [playlist, setPlaylist] = useState<Playlist>();
  const [duration, setDuration] = useState<number>(0);
  const navigate = useNavigate();
  useEffect(() => {
    updateSong();
  }, [id]);

  const updateSong = () => {
    if (id == null) return;
    axios
      .get("http://localhost:4000/auth/playlist-id?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Playlist>>) => {
        const playlist = res.data.data;
        console.log(res.data.data);
        setPlaylist(playlist);

        let minute = 0;
        playlist.playlistDetails.map((detail) => {
          minute += detail.song.duration;
        });
        setDuration(minute);
        console.log(playlist);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  };

  const deletePlaylist = () => {
    if (playlist == null) return;
    if (user == null) return;
    void axios
      .delete(
        "http://localhost:4000/auth/playlist?id=" +
          playlist.playlistId +
          "&userId=" +
          user.user_id,
        {
          withCredentials: true,
        },
      )
      .then((res) => {
        console.log(res);
        updateSong();
        updatePlaylist();
        navigate("/home");
      });
  };

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

  return (
    <div className={"outer"}>
      <div className={"App"}>
        <SideBar />

        <Main setSearch={null}>
          <div className={"profileHeader"}>
            <div>
              <img className={"song"} src={playlist?.image} alt={"avatar"} />
            </div>
            <div style={{ margin: "1rem" }}>
              <p>Playlist</p>
              <h1>{playlist?.title}</h1>
              {playlist && <RichText description={playlist.description} />}
              <div className={"songDescription"}>
                <img
                  src={
                    playlist?.user.avatar
                      ? playlist.user.avatar
                      : "/assets/download (6).png"
                  }
                  alt={"avatar"}
                />{" "}
                - {playlist?.user.username} - {playlist?.playlistDetails.length}{" "}
                songs - {Math.floor(duration / 60)} min{" "}
                {Math.floor(duration % 60)
                  .toString()
                  .padStart(2, "0")}{" "}
                sec
              </div>
            </div>
          </div>
          <div className={"artistPlay"}>
            <div className={"playWrapper"}>
              <Play />
            </div>
            {user && user.user_id === playlist?.user.user_id && (
              <Trash2 onClick={deletePlaylist} />
            )}
            {/*<FollowButton userFollow={song.user}/>*/}
          </div>
          <div className={"albumSong"}>
            <div className="cardWrapper">
              <div className={"playlistTable"}>
                <div className={"title"}>
                  <h3>#</h3>
                  <p className={"head"}>Title</p>
                </div>
                <p>Album</p>
                <p>Date added</p>
                <Clock />
                <p></p>
              </div>
              <hr />
            </div>
            <div className="cardWrapper">
              {playlist?.playlistDetails.map((detail, index) => (
                <PlaylistTable
                  updateSong={updateSong}
                  currUser={playlist.user}
                  detail={detail}
                  index={index}
                  key={detail.playlistDetailId}
                />
              ))}
            </div>
          </div>
          <Footer />
        </Main>
        <RightSideBar />
      </div>
      <ControlMusic />
    </div>
  );
};
