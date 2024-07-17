import { useNavigate } from "react-router-dom";

export const SidePlaylist = ({ playlist }: { playlist: Playlist }) => {
  const navigate = useNavigate();
  return (
    <div
      className="sidePlaylist"
      onClick={() => {
        navigate("/playlist/" + playlist.playlistId);
      }}
      style={{ cursor: "pointer" }}
      key={playlist.playlistId}
    >
      <img src={playlist.image} alt={""} />
      <div className={"sidePlaylistContent"}>
        <h3>{playlist.title}</h3>
        <p>{playlist.user.username}</p>
      </div>
    </div>
  );
};
