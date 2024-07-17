import type { AxiosError, AxiosResponse } from "axios";
import axios from "axios";
import { Pencil } from "lucide-react";
import { type ChangeEvent, useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";

import { Card } from "../component/Card.tsx";
import { ControlMusic } from "../component/ControlMusic.tsx";
import { ErrorModal } from "../component/ErrorModal.tsx";
import { Main } from "../component/Main.tsx";
import { ProfileCard } from "../component/ProfileCard.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const ProfilePage = () => {
  const { user, getUser, authenticated } = useAuth();
  const { id } = useParams<{ id: string }>();
  const [userProfile, setUserProfile] = useState<User>({} as User);
  const [follower, setFollower] = useState<Follow[]>();
  const [mutual, setMutual] = useState<Follow[]>();
  const [following, setFollowing] = useState<Follow[]>();
  const [playlist, setPlaylist] = useState<Playlist[]>();
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");
  const navigate = useNavigate();

  useEffect(() => {
    if (user == null || id == undefined) return;

    axios
      .get("http://localhost:4000/auth/user/get?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<User>>) => {
        setUserProfile(res.data.data);
        if (
          res.data.data.role == "Artist" &&
          user.user_id != res.data.data.user_id
        ) {
          navigate("/artist/" + id);
        }
        console.log(res.data);
      })
      .catch((err: unknown) => {
        console.log(err);
      });

    axios
      .get("http://localhost:4000/auth/playlist?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Playlist[]>>) => {
        setPlaylist(res.data.data);
        console.log(res.data.data);
      })
      .catch((err: unknown) => {
        console.log(err);
      });

    axios
      .get("http://localhost:4000/auth/get-following?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
        setFollowing(res.data.data);
        console.log("following");
        console.log(res.data.data);
      })
      .catch((err: unknown) => {
        console.log(err);
      });

    axios
      .get("http://localhost:4000/auth/get-follower?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
        setFollower(res.data.data);
        console.log("following");
        console.log(res.data.data);
      })
      .catch((err: unknown) => {
        console.log(err);
      });

    axios
      .get("http://localhost:4000/auth/get-mutual?id=" + id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
        setMutual(res.data.data);
        console.log("following");
        console.log(res.data.data);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  }, [id, user]);

  const onEdit = (e: ChangeEvent<HTMLInputElement>) => {
    const image = e.target.files?.[0];
    if (image == null) return;
    if (
      image.name.endsWith(".jpg") &&
      image.name.endsWith(".png") &&
      image.name.endsWith(".jpeg")
    ) {
      setError("Please select an image with jpg, jpeg, or png format");
      return;
    }
    if (user == null) return;
    const formData = new FormData();
    formData.append("image", image as Blob);
    formData.append("userId", user.user_id);
    axios
      .post("http://localhost:4000/auth/user/update-pic", formData, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<string>>) => {
        console.log(res);
        setSuccess("Profile Updated SuccessFully!");
        getUser();
      })
      .catch((err: unknown) => {
        const error = err as AxiosError<WebResponse<string>>;
        if (error.response == undefined) return;
        setError(error.response.data.message);
      });
  };

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

  return (
    <div className={"outer"}>
      <div className={"App"}>
        <SideBar />
        {error && <ErrorModal error={error} setError={setError} />}
        {success && <SuccessModal success={success} setSuccess={setSuccess} />}
        <Main setSearch={null}>
          <div className={"profileHeader"}>
            <div>
              {userProfile.user_id == user?.user_id ? (
                <label htmlFor={"image"}>
                  <img
                    src={
                      userProfile.avatar
                        ? userProfile.avatar
                        : "/assets/download (6).png"
                    }
                    alt={"avatar"}
                  />
                  <Pencil />
                </label>
              ) : (
                <img
                  src={
                    userProfile.avatar
                      ? userProfile.avatar
                      : "/assets/download (6).png"
                  }
                  alt={"avatar"}
                />
              )}
            </div>
            <div>
              <p>Profile</p>
              <h1>{userProfile.username}</h1>
              <h6>
                {playlist?.length} Public Playlists - {follower?.length}{" "}
                Followers - {following?.length} Following
              </h6>
            </div>
            {/*<FollowButton userFollow={userProfile} />*/}
          </div>
          {id && playlist && playlist.length > 0 && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Public Playlists</h2>
                <Link to={"/more/playlist/public/" + id}>Show More..</Link>
              </div>
              <div className="cardWrapper">
                {playlist.slice(0, 5).map((play) => (
                  <Card playlist={play} key={play.playlistId} play={false} />
                ))}
              </div>
            </div>
          )}
          {id && follower && follower.length > 0 && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Followers</h2>
                <Link to={"/more/profile/follower/" + id}>Show More..</Link>
              </div>
              <div className="cardWrapper">
                {follower.slice(0, 5).map((follow) => (
                  <ProfileCard
                    user={follow.Follower}
                    key={follow.Follower.user_id}
                  />
                ))}
              </div>
            </div>
          )}
          {id && following && following.length > 0 && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Following</h2>
                <Link to={"/more/profile/following/" + id}>Show More..</Link>
              </div>
              <div className="cardWrapper">
                {following.slice(0, 5).map((follow) => (
                  <ProfileCard
                    user={follow.Following}
                    key={follow.Following.user_id}
                  />
                ))}
              </div>
            </div>
          )}
          {id && mutual && mutual.length > 0 && (
            <div className="cardContainer">
              <div className={"cardTitle"}>
                <h2>Mutual</h2>
                <Link to={"/more/profile/mutual/" + id}>Show More..</Link>
              </div>
              <div className="cardWrapper">
                {mutual.slice(0, 5).map((follow) => (
                  <ProfileCard
                    user={follow.Follower}
                    key={follow.Follower.user_id}
                  />
                ))}
              </div>
            </div>
          )}
          <input
            type={"file"}
            id={"image"}
            name={"image"}
            style={{ opacity: 0, display: "none" }}
            onChange={onEdit}
          />
        </Main>
        <RightSideBar />
      </div>
      <ControlMusic />
    </div>
  );
};
