import type { AxiosResponse } from "axios";
import axios from "axios";
import { Plus } from "lucide-react";
import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import { AlbumCard } from "../component/AlbumCard.tsx";
import { ControlMusic } from "../component/ControlMusic.tsx";
import { Main } from "../component/Main.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const YourPostPage = () => {
  const navigate = useNavigate();
  const { user } = useAuth();
  const [artist, setArtist] = useState<Artist | null>(null);
  const [albums, setAlbum] = useState<Album[]>();
  useEffect(() => {
    if (user == null) return;
    if (user.role != "Artist") navigate("/home");

    axios
      .get("http://localhost:4000/auth/artist/get?id=" + user.user_id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Artist>>) => {
        const artists = res.data.data;
        setArtist(artists);
        axios
          .get(
            "http://localhost:4000/auth/album/get-artist?id=" +
              artists.artistId,
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
  }, [user]);
  return (
    <div className={"outer"}>
      <div className={"App"}>
        <SideBar />

        <Main setSearch={null}>
          <div className="profileHeader">
            <div>
              <img src={artist?.banner} alt={"avatar"} />
            </div>
            <div>
              <p>Verified Artist</p>
              <h1>Hi, {artist?.user.username}</h1>
              {/*<h6>{playlist.length} Public Playlists - {follower.length} Followers*/}
              {/*    - {following.length} Following</h6>*/}
            </div>
          </div>
          <div className="cardContainer">
            <div className={"cardTitle"}>
              <h2>Discography</h2>
              <Link to={"/more/"}>See discography</Link>
            </div>
            <div className="cardWrapper">
              <div
                className={"card"}
                onClick={() => {
                  navigate("/create/music/");
                }}
              >
                <div className={"plusImage"}>
                  <Plus />
                </div>
                <div className={"cardContent"}></div>
              </div>
              {albums?.map((album) => (
                <AlbumCard album={album} key={album.albumId} play={false} />
              ))}
            </div>
          </div>
        </Main>
        <RightSideBar />
      </div>
      <ControlMusic />
    </div>
  );
};
