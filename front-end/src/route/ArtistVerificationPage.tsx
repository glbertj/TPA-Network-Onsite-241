import type { AxiosResponse } from "axios";
import axios from "axios";
import { Check, X } from "lucide-react";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const ArtistVerificationPage = () => {
  const { user, authenticated } = useAuth();
  const [error, setError] = useState<string>("");
  const [artist, setArtist] = useState<VerifyArtist[] | null>(null);

  const onVerify = (artistId: string) => {
    axios
      .put(
        "http://localhost:4000/admin/artist/update?id=" + artistId,
        {},
        {
          withCredentials: true,
        },
      )
      .then((res) => {
        getUnverified();
        console.log(res);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  };

  const onCancel = (artistId: string, userId: string) => {
    if (user == null) return;
    axios
      .delete(
        "http://localhost:4000/admin/artist/delete?id=" +
          artistId +
          "&userId=" +
          userId,
        {
          withCredentials: true,
        },
      )
      .then((res) => {
        getUnverified();
        console.log(res);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  };

  const navigate = useNavigate();

  const getUnverified = () => {
    if (user == null) return;
    if (user.role != "Admin") navigate("/home");

    axios
      .get("http://localhost:4000/auth/artist/get-unverified", {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Artist[] | null>>) => {
        const artist = res.data.data;
        if (artist == null) {
          setArtist(null);
          return;
        }
        const verifyArtist: VerifyArtist[] = [];
        artist.map((art) => {
          axios
            .get("http://localhost:4000/auth/get-following?id=" + art.userId, {
              withCredentials: true,
            })
            .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
              console.log(res.data.data);
            })
            .catch((err: unknown) => {
              console.log(err);
            });

          axios
            .get("http://localhost:4000/get-follower?id=" + art.userId, {
              withCredentials: true,
            })
            .then((res: AxiosResponse<WebResponse<Follow[]>>) => {
              console.log(res.data.data);
            })
            .catch((err: unknown) => {
              console.log(err);
            });
          verifyArtist.push({ artist: art, follower: 0, following: 0 });
        });
        setArtist(verifyArtist);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  };

  useEffect(() => {
    getUnverified();
  }, [user]);

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
    if (authenticated) {
      if (user?.role != "Admin") navigate("/home");
    }
  }, [authenticated]);

  return (
    <div className={"wrapper"}>
      {error && <ErrorModal error={error} setError={setError} />}
      <Navbar />
      <div className="container">
        <div className={"adminContainer"}>
          <div className={"adminBox"}>
            <div className={"editProfileTitle"}>
              <h1>Admin Page</h1>
              <p>Verify Artist</p>
            </div>

            {artist == null && <h1>No artist to verify</h1>}
            {artist?.map((art) => {
              return (
                <div className={"verifyContainer"} key={art.artist.artistId}>
                  <div className={"left"}>
                    <div>
                      <img
                        src={art.artist.banner}
                        alt={"art"}
                        onClick={() => {
                          const url = "/profile/" + art.artist.userId;
                          window.open(url, "_blank");
                          // navigate("/profile/" + art.artist.userId);
                        }}
                        style={{ cursor: "pointer" }}
                      />
                    </div>
                    <div className={"userContent"}>
                      <h6>{art.artist.user.username}</h6>
                      <p>
                        {art.follower} Follower - {art.following} Following
                      </p>
                    </div>
                  </div>
                  <div className={"right"}>
                    <div
                      onClick={() => {
                        onVerify(art.artist.artistId);
                      }}
                    >
                      <Check className={"check"} />
                    </div>
                    <div
                      onClick={() => {
                        onCancel(art.artist.artistId, art.artist.userId);
                      }}
                    >
                      <X className={"x"} />
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};
