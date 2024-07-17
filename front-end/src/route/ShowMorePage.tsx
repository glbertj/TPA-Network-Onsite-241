import type { AxiosResponse } from "axios";
import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import { AlbumCard } from "../component/AlbumCard.tsx";
import { Card } from "../component/Card.tsx";
import { ControlMusic } from "../component/ControlMusic.tsx";
import { Main } from "../component/Main.tsx";
import { ProfileCard } from "../component/ProfileCard.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { AlbumSkeleton } from "../component/skeleton/AlbumSkeleton.tsx";
import { SongCard } from "../component/SongCard.tsx";

export const ShowMorePage = () => {
  const { type, subtype, id } = useParams();
  const [recommendation, setRecommendation] = useState<Album[]>();
  const [isLoad, setIsLoad] = useState<boolean>(false);
  const [page, setPage] = useState(2);
  const [gallery, setGallery] = useState<Play[]>();
  const [follower, setFollower] = useState<Follow[]>();
  const [mutual, setMutual] = useState<Follow[]>();
  const [following, setFollowing] = useState<Follow[]>();
  const [playlist, setPlaylist] = useState<Playlist[]>();
  const [songs, setSongs] = useState<Song[]>();

  useEffect(() => {
    if (subtype == undefined || id == undefined) return;
  }, [subtype, id]);

  useEffect(() => {
    if (subtype == undefined || id == undefined) return;
    console.log(subtype);
    if (subtype == "browse") {
      axios
        .get("http://localhost:4000/auth/song/get-all", {
          withCredentials: true,
        })
        .then((res: AxiosResponse<WebResponse<Song[]>>) => {
          setSongs(res.data.data);
          console.log(res);
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }

    if (subtype == "recommendation") {
      const fetchRecommendations = () => {
        setIsLoad(true);
        setTimeout(() => {
          axios
            .get("http://localhost:4000/auth/album/get-random", {
              withCredentials: true,
            })
            .then((res: AxiosResponse<WebResponse<Album[]>>) => {
              setRecommendation((prev) => {
                if (prev == undefined) return res.data.data;
                return [...prev, ...res.data.data];
              });
              setIsLoad(false);
            })
            .catch((err: unknown) => {
              console.log(err);
            });
        }, 1000);
      };
      fetchRecommendations();
    }
    if (subtype == "recently") {
      axios
        .get("http://localhost:4000/auth/play/get-last-rec?id=" + id, {
          withCredentials: true,
        })
        .then((res: AxiosResponse<WebResponse<Play[]>>) => {
          setGallery(res.data.data);
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }

    if (subtype == "public") {
      axios
        .get("http://localhost:4000/auth/playlist?id=" + id, {
          withCredentials: true,
        })
        .then((res: AxiosResponse<WebResponse<Playlist[]>>) => {
          setPlaylist(res.data.data);
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }

    if (subtype == "following") {
      axios
        .get("http://localhost:4000/auth/get-following?id=" + id, {
          withCredentials: true,
        })
        .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
          setFollowing(res.data.data);
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }

    if (subtype == "follower") {
      axios
        .get("http://localhost:4000/auth/get-follower?id=" + id, {
          withCredentials: true,
        })
        .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
          setFollower(res.data.data);
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }

    if (subtype == "mutual") {
      axios
        .get("http://localhost:4000/auth/get-mutual?id=" + id, {
          withCredentials: true,
        })
        .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
          setMutual(res.data.data);
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }
  }, [subtype, page, id]);

  const handleScroll = () => {
    if (subtype != "recommendation") return;
    const content = document.getElementById("content") as HTMLDivElement;

    const { scrollTop, scrollHeight, clientHeight } = content;
    if (scrollTop + clientHeight + 1 >= scrollHeight) {
      setPage((prev) => prev + 5);
    }
  };

  useEffect(() => {
    if (subtype != "recommendation") return;
    document
      .getElementById("content")
      ?.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <div className={"outer"}>
      <div className={"App"}>
        <SideBar />

        <Main setSearch={null}>
          <div className={"profileHeader"}>
            <div>
              <h1>
                {type?.substring(0, 1).toUpperCase().concat(type.substring(1))}
              </h1>
            </div>
          </div>
          {subtype === "browse" && songs && (
            <div className="cardContainer">
              <div className="cardTitle">
                <h2>Browse Music</h2>
              </div>
              <div className="cardWrapper">
                {songs.length > 0 ? (
                  songs.map((song, index) => (
                    <SongCard key={index} songs={song} play={true} />
                  ))
                ) : (
                  <div>No recommendations available</div>
                )}
              </div>
            </div>
          )}
          {subtype === "recommendation" && recommendation && (
            <div className="cardContainer">
              <div className="cardTitle">
                <h2>Recommendation</h2>
              </div>
              <div className="cardWrapper">
                {recommendation.length > 0 ? (
                  recommendation.map((album, index) => (
                    <AlbumCard key={index} album={album} play={false} />
                  ))
                ) : (
                  <div>No recommendations available</div>
                )}
              </div>
            </div>
          )}
          {subtype == "recommendation" && isLoad && (
            <div className="cardContainer">
              <div className="cardWrapper">
                {Array(5)
                  .fill(0)
                  .map((_, index) => (
                    <AlbumSkeleton key={index} />
                  ))}
              </div>
            </div>
          )}
          {subtype == "recently" && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Recently Played</h2>
              </div>
              <div className="cardWrapper">
                {gallery?.map((play) => (
                  <AlbumCard
                    album={play.song.album}
                    key={play.playId}
                    play={false}
                  />
                ))}
              </div>
            </div>
          )}
          {subtype == "public" && playlist && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Public Playlists</h2>
              </div>
              <div className="cardWrapper">
                {playlist.map((play) => (
                  <Card playlist={play} key={play.playlistId} play={false} />
                ))}
              </div>
            </div>
          )}
          {subtype == "follower" && follower && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Followers</h2>
              </div>
              <div className="cardWrapper">
                {follower.map((follow) => (
                  <ProfileCard
                    user={follow.Follower}
                    key={follow.Follower.user_id}
                  />
                ))}
              </div>
            </div>
          )}
          {subtype == "following" && following && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Following</h2>
              </div>
              <div className="cardWrapper">
                {following.map((follow) => (
                  <ProfileCard
                    user={follow.Following}
                    key={follow.Following.user_id}
                  />
                ))}
              </div>
            </div>
          )}
          {subtype == "mutual" && mutual && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Mutual</h2>
              </div>
              <div className="cardWrapper">
                {mutual.map((follow) => (
                  <ProfileCard
                    user={follow.Follower}
                    key={follow.Follower.user_id}
                  />
                ))}
              </div>
            </div>
          )}
        </Main>
        <RightSideBar />
      </div>
      <ControlMusic />
    </div>
  );
};
