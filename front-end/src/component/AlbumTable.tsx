import React from "react";
import { useNavigate } from "react-router-dom";

export const AlbumTable = ({ song, index }: { song: Song; index: number }) => {
  const navigate = useNavigate();
  const handleImageClick = (e: React.MouseEvent<HTMLSpanElement>) => {
    e.stopPropagation();
    navigate("/album/" + song.albumId);
  };
  return (
    <div
      className={"albumTable"}
      onClick={() => {
        navigate("/track/" + song.songId);
      }}
    >
      <div className={"title"}>
        <p>{index + 1}. </p>
        <img
          src={song.album.banner}
          alt="Song Cover"
          onClick={handleImageClick}
        />
        <div>
          <h3>{song.title}</h3>
          <p>{song.album.title}</p>
        </div>
      </div>
      <p>{song.play.length}</p>
      <p>
        {Math.floor(song.duration / 60)}:
        {Math.floor(song.duration % 60)
          .toString()
          .padStart(2, "0")}
      </p>
    </div>
  );
};
