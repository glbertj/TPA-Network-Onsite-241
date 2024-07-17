import { Play } from "lucide-react";
import React from "react";
import { Link } from "react-router-dom";

export const AlbumCard = ({ album, play }: { album: Album; play: boolean }) => {
  const handlePlayClick = (e: React.MouseEvent<HTMLSpanElement>) => {
    e.stopPropagation();
  };

  return (
    <div className={"card"}>
      <Link to={"/album/" + album.albumId}>
        <div className={"cardImage"}>
          <img src={album.banner} alt={"placeholder"} />
          {play && (
            <span className={"play"} onClick={handlePlayClick}>
              <Play />
            </span>
          )}
        </div>
        <div className={"cardContent"}>
          <h3>{album.title}</h3>
          <p>
            {new Date(album.release).getFullYear()} - {album.type}
          </p>
        </div>
      </Link>
    </div>
  );
};
