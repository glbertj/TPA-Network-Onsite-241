import { BadgeCheck, X } from "lucide-react";

import { useSong } from "../context/UseSong.tsx";
import { Advertisement } from "./Advertisement.tsx";
import { FollowButton } from "./FollowButton.tsx";
import { Queue } from "./Queue.tsx";
import { SideSong } from "./SideSong.tsx";

export const RightSideBar = () => {
  const { showDetail, showDetailHandler, song, track, waitingSong } = useSong();

  return (
    <>
      {showDetail && (
        <div className="rightSideBar">
          {showDetail === "detail" ? (
            <>
              <div className="rightSideBarHeader">
                <h3>{track ? track : song?.title}</h3>
                <X
                  onClick={() => {
                    showDetailHandler("");
                  }}
                />
              </div>
              <div className="trackImage">
                <img src={song?.album.banner} alt="Song Cover" />
              </div>
              <div className="songTitle">
                <div>
                  <h3>{song?.title}</h3>
                  <p>{song?.artist.user.username}</p>
                </div>
                <div>
                  <BadgeCheck />
                </div>
              </div>
              <div className="aboutArtist">
                <div className="header">
                  <h3>About the Artist</h3>
                  <img
                    src={
                      song?.artist.user.avatar
                        ? song.artist.user.avatar
                        : "/assets/download (6).png"
                    }
                    alt="Artist Avatar"
                  />
                </div>
                <div className="aboutContent">
                  <h3>{song?.artist.user.username}</h3>
                  <div className="description">
                    <p>{song?.play?.length} monthly listeners</p>
                    {song && <FollowButton userFollow={song.artist.user} />}
                  </div>
                  <p>{song?.artist.description}</p>
                </div>
              </div>
              <div className="queue">
                <div className="header">
                  <h3>Next in queue</h3>
                  <button
                    onClick={() => {
                      showDetailHandler("queue");
                    }}
                  >
                    Open queue
                  </button>
                </div>
                {waitingSong ? (
                  <SideSong songs={waitingSong[0]} trash={false} index={0} />
                ) : (
                  <div className="sideSong">
                    <div className={"albumPic"}>
                      <img
                        src={"/assets/download (6).png"}
                        alt={"songs.title"}
                        className="albumPic"
                      />
                    </div>
                    <div className="song-details">
                      <h3 className="song-title">No Song in Queue</h3>
                    </div>
                  </div>
                )}
              </div>
            </>
          ) : showDetail === "advertise" ? (
            <Advertisement />
          ) : (
            <Queue />
          )}
        </div>
      )}
    </>
  );
};
