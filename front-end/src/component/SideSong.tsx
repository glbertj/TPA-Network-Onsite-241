import { Trash2 } from "lucide-react";
import { useNavigate } from "react-router-dom";

import { useAuth } from "../context/UseAuth.tsx";
import { useSong } from "../context/UseSong.tsx";

export const SideSong = ({
  songs,
  trash,
  index,
}: {
  songs: Song;
  trash: boolean;
  index: number;
}) => {
  const { song, removeQueue, dequeue, waitingSong, setSong } = useSong();
  const { user } = useAuth();
  const navigate = useNavigate();
  const handleNavigate = () => {
    navigate("/track/" + songs.songId);
  };

  const handleRemoveQueue = (e: React.MouseEvent<SVGSVGElement>) => {
    e.stopPropagation();
    if (waitingSong == null) {
      setSong(null);
    } else if (songs.songId == song?.songId) {
      dequeue(user);
    } else {
      removeQueue(index, user);
    }
  };

  return (
    <div className="sideSong" onClick={handleNavigate}>
      <div className={"albumPic"}>
        {/*{song != songs && <Play />}*/}
        <img src={songs.album.banner} alt={songs.title} className="albumPic" />
      </div>
      <div className="song-details">
        <h3 className="song-title">{songs.title}</h3>
        <p className="artist-name">{songs.artist.user.username}</p>
      </div>
      <div className={"trash"}>
        {trash && <Trash2 onClick={handleRemoveQueue} />}
      </div>
    </div>
  );
};
