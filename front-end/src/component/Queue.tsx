import { X } from "lucide-react";

import { useSong } from "../context/UseSong.tsx";
import { SideSong } from "./SideSong.tsx";

export const Queue = () => {
  const { showDetailHandler, song, track, waitingSong, clearAllQueue } =
    useSong();

  return (
    <>
      <div className="rightSideBarHeader">
        <h3>{track ? track : song?.title}</h3>
        <X
          onClick={() => {
            showDetailHandler("");
          }}
        />
      </div>
      {song && (
        <div className="queue">
          <div className="header">
            <h3>Now playing</h3>
            <button onClick={clearAllQueue}>Clear queue</button>
          </div>
          <SideSong songs={song} trash={true} index={0} />
        </div>
      )}
      <div className="queue">
        <div className="header">
          {!waitingSong && <h3>No song in queue</h3>}
          {waitingSong && waitingSong.length <= 0 && <h3>No song in queue</h3>}
          {waitingSong && waitingSong.length > 0 && <h3>Next in queue</h3>}
        </div>
        {waitingSong?.slice(0, waitingSong.length).map((song, index) => {
          return (
            <SideSong
              songs={song}
              trash={true}
              key={song.songId}
              index={index}
            />
          );
        })}
      </div>
    </>
  );
};
