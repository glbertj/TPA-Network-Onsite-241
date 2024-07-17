import type { AxiosResponse } from "axios";
import axios from "axios";
import { useEffect, useState } from "react";

import { useAuth } from "../context/UseAuth.tsx";

export const FollowButton = ({ userFollow }: { userFollow: User }) => {
  const { user } = useAuth();
  const [isFollowing, setIsFollowing] = useState(false);

  const GetButton = () => {
    if (user) {
      axios
        .get("http://localhost:4000/auth/get-following?id=" + user.user_id, {
          withCredentials: true,
        })
        .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
          if (
            res.data.data.find(
              (follow) => follow.FollowingId === userFollow.user_id,
            )
          ) {
            setIsFollowing(true);
          } else {
            setIsFollowing(false);
          }
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }
  };

  useEffect(() => {
    GetButton();
  }, [user]);

  const handleFollow = () => {
    if (user) {
      if (user.user_id == userFollow.user_id) return;
      if (isFollowing) {
        axios
          .delete("http://localhost:4000/auth/follow", {
            data: {
              followId: userFollow.user_id,
              followerId: user.user_id,
            },
            withCredentials: true,
          })
          .then((res) => {
            console.log(res);
            // setIsFollowing(false);
            GetButton();
          })
          .catch((err: unknown) => {
            console.log(err);
          });
      } else {
        axios
          .put(
            "http://localhost:4000/auth/follow",
            {
              followId: userFollow.user_id,
              followerId: user.user_id,
            },
            {
              withCredentials: true,
            },
          )
          .then((res) => {
            console.log(res);
            // setIsFollowing(true);
            GetButton();
          })
          .catch((err: unknown) => {
            console.log(err);
          });
      }
    }
  };
  return (
    <div>
      {user && user.user_id != userFollow.user_id && (
        <button className={"follow"} onClick={handleFollow}>
          {isFollowing ? "Unfollow" : "Follow"}
        </button>
      )}
    </div>
  );
};
