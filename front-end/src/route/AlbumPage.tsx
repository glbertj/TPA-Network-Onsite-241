import type { AxiosResponse } from "axios";
import axios from "axios";
import { Clock, Dot, Play } from "lucide-react";
import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";

import { AlbumCard } from "../component/AlbumCard.tsx";
import { AlbumTable } from "../component/AlbumTable.tsx";
import { ControlMusic } from "../component/ControlMusic.tsx";
import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Main } from "../component/Main.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { useSong } from "../context/UseSong.tsx";

export const AlbumPage = () => {
  const { id } = useParams<{ id: string }>();
  const { authenticated, user } = useAuth();
  const { clearAllQueue, enqueue } = useSong();
  const navigate = useNavigate();
  const [error, setError] = useState<string>("");
  const [album, setAlbum] = useState<Album>();
  const [duration, setDuration] = useState<number>(0);
  const [albumSong, setAlbumSong] = useState<Song[]>();
  const [moreAlbum, setMoreAlbum] = useState<Album[]>();
  const [songCount, setSongCount] = useState<number>(0);
  useEffect(() => {
    if (id == null) return;
    axios
      .get("http://localhost:4000/auth/song/get-by-album?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Song[]>>) => {
        setAlbumSong(res.data.data);
        console.log(res.data.data);
        const albums: Album = res.data.data[0].album;
        albums.artist = res.data.data[0].artist;
        setAlbum(albums);
        let minute = 0;
        let count = 0;
        res.data.data.map((song) => {
          minute += song.duration;
          count += 1;
        });
        setSongCount(count);
        setDuration(minute);

        axios
          .get(
            "http://localhost:4000/auth/album/get-artist?id=" + albums.artistId,
            {
              withCredentials: true,
            },
          )
          .then((res: AxiosResponse<WebResponse<Album[]>>) => {
            setMoreAlbum(res.data.data);
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
    if (albumSong == null) {
      setError("No song in this album");
      return;
    }
    clearAllQueue();
    albumSong.map((song) => {
      enqueue(song, user);
    });
  };

  return (
    <div className={"outer"}>
      <div className={"App"}>
        <SideBar />
        {error && <ErrorModal error={error} setError={setError} />}
        <Main setSearch={null}>
          <div className={"profileHeader"}>
            <div>
              <img className={"song"} src={album?.banner} alt={"avatar"} />
            </div>
            <div>
              <p>{album?.type}</p>
              <h1>{album?.title}</h1>
              <div className={"songDescription"}>
                <img
                  src={
                    album?.artist.user.avatar
                      ? album.artist.user.avatar
                      : "/assets/download (6).png"
                  }
                  alt={"avatar"}
                />{" "}
                <Dot /> {album?.artist.user.username} <Dot />{" "}
                {album && new Date(album.release).getFullYear()} <Dot />{" "}
                {songCount} songs <Dot /> {Math.floor(duration / 60)} min{" "}
                {Math.floor(duration % 60)
                  .toString()
                  .padStart(2, "0")}{" "}
                sec
              </div>
            </div>
          </div>
          <div className={"artistPlay"}>
            <div className={"playWrapper"} onClick={handlePlay}>
              <Play />
            </div>
            {/*<FollowButton userFollow={song.user}/>*/}
          </div>
          <div className={"albumSong"}>
            <div className="cardWrapper">
              <div className={"albumTable"}>
                <div className={"title"}>
                  <h3>#</h3>
                  <p className={"head"}>Title</p>
                </div>
                <p>Listen Count</p>
                <Clock />
              </div>
              <hr />
            </div>
            <div className="cardWrapper">
              {albumSong?.map((song, index) => (
                <AlbumTable song={song} index={index} key={song.songId} />
              ))}
            </div>
          </div>
          {moreAlbum && moreAlbum.length > 0 && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Discography</h2>
                <Link to={"/more/"}>See discography</Link>
              </div>
              <div className="cardWrapper">
                {moreAlbum.slice(0, 5).map((album) => (
                  <AlbumCard album={album} key={album.albumId} play={false} />
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
