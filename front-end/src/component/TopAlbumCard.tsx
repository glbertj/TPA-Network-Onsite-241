import { Play } from "lucide-react";
import React from "react";

export const TopAlbumCard = ({
  result,
  handleNavigate,
  play,
}: {
  result: SearchResponse;
  handleNavigate: (type: string, result: SearchResponse) => void;
  play: boolean;
}) => {
  const handlePlayClick = (e: React.MouseEvent<HTMLSpanElement>) => {
    e.stopPropagation();
  };

  return (
    <div
      className={"card"}
      onClick={() => {
        handleNavigate("album", result);
      }}
    >
      <div className={"cardImage"}>
        <img src={result.song.album.banner} alt={"placeholder"} />
        {play && (
          <span className={"play"} onClick={handlePlayClick}>
            <Play />
          </span>
        )}
      </div>
      <div className={"cardContent"}>
        <h3>{result.song.album.title}</h3>
        <p>
          {new Date(result.song.album.release).getFullYear()} -{" "}
          {result.song.album.type}
        </p>
      </div>
    </div>
  );
};
