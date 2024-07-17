import type { AxiosResponse } from "axios";
import axios from "axios";
import { Dot, Play } from "lucide-react";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

import { AddPlaylistButton } from "../component/AddPlaylistButton.tsx";
import { AlbumTable } from "../component/AlbumTable.tsx";
import { ControlMusic } from "../component/ControlMusic.tsx";
import { FollowButton } from "../component/FollowButton.tsx";
import { Footer } from "../component/Footer.tsx";
import { Main } from "../component/Main.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { SongTable } from "../component/SongTable.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { useSong } from "../context/UseSong.tsx";

export const TrackPage = () => {
  const { clearAllQueue, enqueue } = useSong();
  const { id } = useParams();
  const { authenticated, user } = useAuth();
  const navigate = useNavigate();
  const [song, setSong] = useState<Song>();
  const [topTrack, setTopTrack] = useState<Song[]>();
  const [albumSong, setAlbumSong] = useState<Song[]>();

  useEffect(() => {
    if (id == null) return;

    axios
      .get("http://localhost:4000/auth/song/get?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Song>>) => {
        const songs = res.data.data;
        setSong(res.data.data);
        console.log(songs);
        axios
          .get(
            "http://localhost:4000/auth/song/get-by-artist?id=" +
              songs.artistId,
            {
              withCredentials: true,
            },
          )
          .then((res: AxiosResponse<WebResponse<Song[]>>) => {
            setTopTrack(res.data.data);
          })
          .catch((err: unknown) => {
            console.log(err);
          });
        axios
          .get(
            "http://localhost:4000/auth/song/get-by-album?id=" + songs.albumId,
            {
              withCredentials: true,
            },
          )
          .then((res: AxiosResponse<WebResponse<Song[]>>) => {
            // console.log(res)
            setAlbumSong(res.data.data);
          })
          .catch((err: unknown) => {
            console.log(err);
          });
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  }, [id]);

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

  const handlePlay = () => {
    if (song == null) return;
    clearAllQueue();
    enqueue(song, user);
  };

  return (
    <div className={"outer"}>
      <div className={"App"}>
        <SideBar />

        <Main setSearch={null}>
          <div className={"profileHeader"}>
            <div>
              <img className={"song"} src={song?.album.banner} alt={"avatar"} />
            </div>
            <div>
              <p>Song</p>
              <h1>{song?.title}</h1>
              <div className={"songDescription"}>
                <img src={song?.artist.banner} alt={"avatar"} /> <Dot />
                {song?.album.title} <Dot />
                {song?.artist.user.username} <Dot />
                {song && new Date(song.releaseDate).getFullYear()} <Dot />
                {song && Math.floor(song.duration / 60)}:
                {song &&
                  Math.floor(song.duration % 60)
                    .toString()
                    .padStart(2, "0")}
                <Dot /> {song?.play.length} plays
              </div>
            </div>
          </div>
          <div className={"artistPlay"}>
            <div className={"playWrapper"} onClick={handlePlay}>
              <Play />
            </div>
            {song && <FollowButton userFollow={song.artist.user} />}
            {song && <AddPlaylistButton song={song} />}
          </div>
          {topTrack && (
            <div className={"popular"}>
              <p>Popular tracks by</p>
              <h2>{song?.artist.user.username}</h2>
              <div className="cardWrapper">
                {topTrack.slice(0, 5).map((song, index) => (
                  <SongTable song={song} index={index} key={song.songId} />
                ))}
              </div>
            </div>
          )}
          {song && (
            <div className={"fromAlbum"}>
              <img src={song.album.banner} alt={"album image"} />
              <div>
                <p>From the {song.album.type}</p>
                <h2>{song.album.title}</h2>
              </div>
            </div>
          )}
          {albumSong && (
            <div className={"albumSong"}>
              <div className="cardWrapper">
                {albumSong.slice(0, 5).map((song, index) => (
                  <AlbumTable song={song} index={index} key={song.songId} />
                ))}
              </div>
            </div>
          )}
          <Footer />
        </Main>
        <RightSideBar />
      </div>
      <ControlMusic />
    </div>
  );
};
