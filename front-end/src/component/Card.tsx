import { Play } from "lucide-react";
import React from "react";
import { useNavigate } from "react-router-dom";

export const Card = ({
  playlist,
  play,
}: {
  playlist: Playlist;
  play: boolean;
}) => {
  const handlePlayClick = (e: React.MouseEvent<HTMLSpanElement>) => {
    e.stopPropagation();
  };

  const navigate = useNavigate();
  console.log(playlist);

  return (
    <div
      className={"card"}
      onClick={() => {
        navigate("/playlist/" + playlist.playlistId);
      }}
    >
      <div className={"cardImage"}>
        <img src={playlist.image} alt={"placeholder"} />
        {play && (
          <span className={"play"} onClick={handlePlayClick}>
            <Play />
          </span>
        )}
      </div>
      <div className={"cardContent"}>
        <h3>{playlist.title}</h3>
        <p>{playlist.user.username}</p>
      </div>
    </div>
  );
};
