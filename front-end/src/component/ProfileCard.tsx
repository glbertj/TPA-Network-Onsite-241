import { useNavigate } from "react-router-dom";

export const ProfileCard = ({ user }: { user: User }) => {
  const navigate = useNavigate();

  const handleNavigates = () => {
    if (user.role == "Artist") {
      navigate("/artist/" + user.user_id);
    }
    if (user.role == "Listener") {
      navigate("/profile/" + user.user_id);
    }
  };

  return (
    <div className={"card"} onClick={handleNavigates}>
      <div className={"cardImage"}>
        <img
          src={user.avatar ? user.avatar : "/assets/download (6).png"}
          alt={"placeholder"}
          className={"profilePic"}
        />
      </div>
      <div className={"cardContent"}>
        <h3>{user.username}</h3>
        <p>{user.role}</p>
      </div>
    </div>
  );
};
