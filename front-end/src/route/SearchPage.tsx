import type { AxiosResponse } from "axios";
import axios from "axios";
import { Dot } from "lucide-react";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import { AlbumCard } from "../component/AlbumCard.tsx";
import { ControlMusic } from "../component/ControlMusic.tsx";
import { Main } from "../component/Main.tsx";
import { ProfileCard } from "../component/ProfileCard.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { AlbumSkeleton } from "../component/skeleton/AlbumSkeleton.tsx";
import { ProfileSkeleton } from "../component/skeleton/ProfileSkeleton.tsx";
import { TableSkeleton } from "../component/skeleton/TableSkeleton.tsx";
import { TopResultSkeleton } from "../component/skeleton/TopResultSkeleton.tsx";
import { SongCard } from "../component/SongCard.tsx";
import { TopAlbumCard } from "../component/TopAlbumCard.tsx";
import { TopProfileCard } from "../component/TopProfileCard.tsx";
import { TopResultTable } from "../component/TopResultTable.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { useDebounce, useQuery } from "../hooks/hooks.ts";

export const SearchPage = () => {
  const query = useQuery();
  const [search, setSearch] = useState<string>(() => {
    return query.get("query") ?? "";
  });
  const debounce = useDebounce(search, 1000);
  const [isLoad, setIsLoad] = useState<boolean>(true);
  const [topResult, setTopResult] = useState<SearchResponse[] | null>(null);
  const [recentSearch, setRecentSearch] = useState<SearchResponse[] | null>(
    null,
  );
  const { authenticated } = useAuth();
  const [artists, setArtists] = useState<SearchResponse[] | null>(null);
  const [albums, setAlbums] = useState<SearchResponse[] | null>(null);
  const navigate = useNavigate();
  const handleNavigate = (type: string, result: SearchResponse) => {
    if (type !== "artist" && type !== "album" && type !== "song") return;
    const storedRecentSearch: string | null =
      localStorage.getItem("recentSearch");
    let recentSearches: SearchResponse[] = [];
    if (storedRecentSearch) {
      try {
        recentSearches = JSON.parse(storedRecentSearch) as SearchResponse[];
      } catch (error) {
        console.error(
          "Failed to parse recent searches from localStorage:",
          error,
        );
        recentSearches = [];
      }
    }
    const response = result;
    response.type = type;
    const exist = recentSearches.find(
      (search) =>
        search.song.songId === result.song.songId &&
        search.type === result.type,
    );
    if (!exist) {
      recentSearches.push(response);
      if (recentSearches.length > 5) {
        recentSearches = recentSearches.slice(1, 6);
      } else {
        recentSearches = recentSearches.slice(0, 5);
      }
      localStorage.setItem("recentSearch", JSON.stringify(recentSearches));
    }
    if (type === "artist") {
      navigate("/artist/" + result.song.artist.userId);
    } else if (type === "album") {
      navigate("/album/" + result.song.albumId);
    } else {
      navigate("/track/" + result.song.songId);
    }
  };

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

  useEffect(() => {
    if (debounce === "") return;
    setIsLoad(true);
    setArtists(null);
    setAlbums(null);
    setTimeout(() => {
      axios
        .get(
          "http://localhost:4000/auth/search/get?keyword=" +
            debounce.toUpperCase(),
          {
            withCredentials: true,
          },
        )
        .then((res: AxiosResponse<WebResponse<SearchResponse[]>>) => {
          setTopResult(res.data.data);
          setIsLoad(false);

          const artistIds = new Set<string>();
          const filterArtists: SearchResponse[] = [];
          res.data.data.forEach((data) => {
            if (!artistIds.has(data.song.artistId)) {
              artistIds.add(data.song.artistId);
              filterArtists.push(data);
            }
          });
          setArtists(filterArtists);

          const albumIds = new Set<string>();
          const filterAlbums: SearchResponse[] = [];
          res.data.data.forEach((data) => {
            if (!albumIds.has(data.song.albumId)) {
              albumIds.add(data.song.albumId);
              filterAlbums.push(data);
            }
          });
          setAlbums(filterAlbums);
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }, 2000);
  }, [debounce]);

  useEffect(() => {
    const storedRecentSearch: string | null =
      localStorage.getItem("recentSearch");
    if (storedRecentSearch === null) return;
    try {
      setRecentSearch(JSON.parse(storedRecentSearch) as SearchResponse[]);
    } catch (error) {
      console.error(
        "Failed to parse recent searches from localStorage:",
        error,
      );
      setRecentSearch(null);
    }
  }, []);

  return (
    <div className="outer">
      <div className="App">
        <SideBar />
        <Main setSearch={setSearch} search={search} setIsLoad={setIsLoad}>
          {debounce === "" ? (
            <div>
              <div className="cardContainer">
                {recentSearch && <h2>Recent Search</h2>}
                <div className="cardWrapper">
                  {recentSearch?.map((result) =>
                    result.type === "artist" ? (
                      <ProfileCard
                        user={result.song.artist.user}
                        key={result.song.artistId}
                      />
                    ) : result.type == "album" ? (
                      <AlbumCard
                        album={result.song.album}
                        key={result.song.albumId}
                        play={false}
                      />
                    ) : (
                      <SongCard
                        songs={result.song}
                        key={result.song.songId}
                        play={false}
                      />
                    ),
                  )}
                </div>
              </div>
              <div className="cardContainer">
                <h2>Browse All</h2>
                <div className="cardWrapper">
                  <div
                    className={"browseMusic"}
                    onClick={() => {
                      navigate("/more/song/browse/all");
                    }}
                  >
                    <h2 className={"title"}>Browse Music</h2>
                    <img src={"/assets/browse.jpg"} alt={":D"} />
                  </div>
                </div>
              </div>
            </div>
          ) : (
            <>
              <div className="searchContainer">
                {isLoad && (
                  <div>
                    <h1 style={{ margin: "0 2rem" }}>Top Results</h1>
                    <TopResultSkeleton />
                  </div>
                )}
                {!isLoad && !topResult?.at(0) && (
                  <div>
                    <h1 style={{ margin: "0 2rem" }}>Top Results</h1>
                    <div className="topResult">
                      <div className="topContent">
                        <h2>No Result Is Found</h2>
                      </div>
                    </div>
                  </div>
                )}
                {!isLoad && topResult?.at(0) && (
                  <div>
                    <h1 style={{ margin: "0 2rem" }}>Top Results</h1>
                    <div className="topResult">
                      <div
                        className="topContent"
                        style={{
                          cursor: "pointer",
                        }}
                        onClick={() => {
                          handleNavigate(topResult[0].type, topResult[0]);
                        }}
                      >
                        <img
                          src={
                            topResult.at(0)?.type === "artist"
                              ? topResult.at(0)?.song.artist.banner
                              : topResult.at(0)?.song.album.banner
                          }
                          alt=""
                          className={
                            topResult.at(0)?.type === "artist"
                              ? "artist"
                              : "song"
                          }
                        />
                        <h2>
                          {topResult.at(0)?.type === "artist"
                            ? topResult.at(0)?.song.artist.user.username
                            : topResult.at(0)?.type === "album"
                              ? topResult.at(0)?.song.album.title
                              : topResult.at(0)?.song.title}
                        </h2>
                        <div className="topTitle">
                          <p>
                            {topResult.at(0)?.type === "artist"
                              ? "Artist"
                              : topResult.at(0)?.type === "album"
                                ? "Album"
                                : "Track"}
                          </p>
                          <Dot />
                          <p>
                            {topResult.at(0)?.type === "artist"
                              ? topResult.at(0)?.song.artist.user.email
                              : topResult.at(0)?.type === "album"
                                ? topResult.at(0)?.song.artist.user.username
                                : topResult.at(0)?.song.artist.user.username}
                          </p>
                        </div>
                      </div>
                    </div>
                  </div>
                )}
                <div>
                  <h1>Songs</h1>
                  <div className="topSongContent">
                    {!isLoad &&
                      topResult
                        ?.slice(0, 5)
                        .map((result, index) => (
                          <TopResultTable
                            result={result}
                            index={index}
                            key={result.song.songId}
                            handleNavigate={handleNavigate}
                          />
                        ))}
                    {!isLoad && !topResult && (
                      <div className="topTable">
                        <div className="title">
                          <p>1. </p>
                          <div
                            style={{
                              background: "#a7a7a7",
                              margin: "0 1rem",
                              width: "50px",
                              height: "50px",
                              borderRadius: "6px",
                            }}
                          ></div>
                          <div>
                            <h3>No Result Found</h3>
                          </div>
                        </div>
                        <p>00:00</p>
                      </div>
                    )}
                    {isLoad &&
                      Array(5)
                        .fill(0)
                        .map((_, index: number) => (
                          <TableSkeleton key={index} />
                        ))}
                  </div>
                </div>
              </div>
              <div className="cardContainer">
                {!isLoad && topResult?.at(0) && <h2>Artists</h2>}
                <div className="cardWrapper">
                  {artists?.map((result) => (
                    <TopProfileCard
                      result={result}
                      handleNavigate={handleNavigate}
                      key={result.song.artistId}
                    />
                  ))}
                </div>
              </div>
              <div className="cardContainer">
                <div className="cardWrapper">
                  {isLoad &&
                    Array(5)
                      .fill(0)
                      .map((_, index: number) => (
                        <ProfileSkeleton key={index} />
                      ))}
                </div>
              </div>
              <div className="cardContainer">
                {!isLoad && topResult?.at(0) && <h2>Collections</h2>}
                <div className="cardWrapper">
                  {albums?.map((result) => (
                    <TopAlbumCard
                      result={result}
                      handleNavigate={handleNavigate}
                      key={result.song.albumId}
                      play={false}
                    />
                  ))}
                </div>
              </div>
              <div className="cardContainer">
                <div className="cardWrapper">
                  {isLoad &&
                    Array(5)
                      .fill(0)
                      .map((_, index: number) => <AlbumSkeleton key={index} />)}
                </div>
              </div>
            </>
          )}
        </Main>
        <RightSideBar />
      </div>
      <ControlMusic />
    </div>
  );
};
