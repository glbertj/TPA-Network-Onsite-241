import type { AxiosResponse } from "axios";
import axios from "axios";
import { Play } from "lucide-react";
import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";

import { AlbumCard } from "../component/AlbumCard.tsx";
import { Card } from "../component/Card.tsx";
import { ControlMusic } from "../component/ControlMusic.tsx";
import { ErrorModal } from "../component/ErrorModal.tsx";
import { FollowButton } from "../component/FollowButton.tsx";
import { Footer } from "../component/Footer.tsx";
import { Main } from "../component/Main.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { SongTable } from "../component/SongTable.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { useSong } from "../context/UseSong.tsx";

export const ArtistPage = () => {
  const { user, authenticated } = useAuth();
  const { id } = useParams<{ id: string }>();
  const { clearAllQueue, enqueue } = useSong();
  const [userProfile, setUserProfile] = useState<Artist>();
  const [playlist, setPlaylist] = useState<Playlist[]>();
  const [filteredAlbum, setFilteredAlbum] = useState<Album[]>();
  const [song, setSong] = useState<Song[]>();
  const [album, setAlbum] = useState<Album[]>();
  const [typeFilter, setType] = useState("all");
  const [error, setError] = useState<string>("");
  const navigate = useNavigate();

  const handleFilter = (type: string) => {
    if (type === typeFilter) {
      setType("all");
    } else if (type === "single") {
      setType("single");
      setFilteredAlbum(album?.filter((album) => album.type === "Single"));
    } else if (type === "album") {
      setType("album");
      setFilteredAlbum(album?.filter((album) => album.type === "Albums"));
    } else if (type === "ep") {
      setType("ep");
      setFilteredAlbum(album?.filter((album) => album.type === "Eps"));
    }
  };

  useEffect(() => {
    if (typeFilter === "all") {
      setFilteredAlbum(album);
    } else if (typeFilter === "single") {
      setFilteredAlbum(album?.filter((album) => album.type === "Single"));
    } else if (typeFilter === "album") {
      setFilteredAlbum(album?.filter((album) => album.type === "Albums"));
    } else if (typeFilter === "ep") {
      setFilteredAlbum(album?.filter((album) => album.type === "Eps"));
    }
  }, [typeFilter, album]);

  useEffect(() => {
    if (user == null || id == null) return;

    axios
      .get("http://localhost:4000/auth/artist/get?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Artist>>) => {
        const profile = res.data.data;
        setUserProfile(profile);
        axios
          .get(
            "http://localhost:4000/auth/song/get-by-artist?id=" +
              profile.artistId,
            {
              withCredentials: true,
            },
          )
          .then((res: AxiosResponse<WebResponse<Song[]>>) => {
            setSong(res.data.data);
            console.log(res.data);
          })
          .catch((err: unknown) => {
            console.log(err);
          });

        axios
          .get("http://localhost:4000/auth/playlist?id=" + user.user_id, {
            withCredentials: true,
          })
          .then((res: AxiosResponse<WebResponse<Playlist[]>>) => {
            const playlist = res.data.data.filter((playlist) =>
              playlist.playlistDetails.find(
                (detail) => detail.song.artistId === profile.artistId,
              ),
            );
            setPlaylist(playlist);
          })
          .catch((err: unknown) => {
            console.log(err);
          });

        axios
          .get(
            "http://localhost:4000/auth/album/get-artist?id=" +
              profile.artistId,
            {
              withCredentials: true,
            },
          )
          .then((res: AxiosResponse<WebResponse<Album[]>>) => {
            setAlbum(res.data.data);
            console.log(res.data);
          })
          .catch((err: unknown) => {
            console.log(err);
          });
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  }, [id, user]);

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

  const handlePlay = () => {
    if (song == null) {
      setError("No song to Play");
      return;
    }
    clearAllQueue();
    song.slice(0,5).map((song) => {
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
              <img src={userProfile?.banner} alt={"avatar"} />
            </div>
            <div>
              <p>Profile</p>
              <h1>{userProfile?.user.username}</h1>
            </div>
          </div>
          <div className={"artistPlay"}>
            <div className={"playWrapper"} onClick={handlePlay}>
              <Play />
            </div>
            {userProfile && <FollowButton userFollow={userProfile.user} />}
          </div>
          {song && (
            <div className={"popular"}>
              <h2>Popular</h2>
              <div className="cardWrapper">
                {song.slice(0, 5).map((song, index) => (
                  <SongTable song={song} index={index} key={song.songId} />
                ))}
              </div>
            </div>
          )}
          {album && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Discography</h2>
                <Link to={"/more/"}>Show More..</Link>
              </div>
              <div className={"filter"}>
                <button
                  className={typeFilter == "single" ? "active" : ""}
                  onClick={() => {
                    handleFilter("single");
                  }}
                >
                  Single
                </button>
                <button
                  className={typeFilter == "album" ? "active" : ""}
                  onClick={() => {
                    handleFilter("album");
                  }}
                >
                  Album
                </button>
                <button
                  className={typeFilter == "ep" ? "active" : ""}
                  onClick={() => {
                    handleFilter("ep");
                  }}
                >
                  EPs
                </button>
              </div>
              <div className="cardWrapper">
                {filteredAlbum
                  ?.slice(0, 5)
                  .map((album) => (
                    <AlbumCard album={album} key={album.albumId} play={false} />
                  ))}
              </div>
            </div>
          )}

          {playlist && playlist.length > 0 && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Featured Playlists</h2>
                <Link to={"/more/"}>Show More..</Link>
              </div>
              {playlist.slice(0, 5).map((play, index) => (
                <div className="cardWrapper" key={index}>
                  <Card playlist={play} play={false} />
                </div>
              ))}
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
