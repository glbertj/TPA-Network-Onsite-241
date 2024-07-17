import { useNavigate } from "react-router-dom";

export const SideUser = ({ user }: { user: User }) => {
  const navigate = useNavigate();
  const handleNavigate = () => {
    if (user.role == "Artist") {
      navigate("/artist/" + user.user_id);
    }
    if (user.role == "Listener") {
      navigate("/profile/" + user.user_id);
    }
  };
  return (
    <div
      className="sideUser"
      onClick={handleNavigate}
      style={{ cursor: "pointer" }}
    >
      <img
        src={user.avatar ? user.avatar : "/assets/download (6).png"}
        alt={""}
      />
      <div className={"sidePlaylistContent"}>
        <h3>{user.username}</h3>
        <p>{user.role}</p>
      </div>
    </div>
  );
};
