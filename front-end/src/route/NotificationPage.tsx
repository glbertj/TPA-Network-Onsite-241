import axios, { type AxiosError } from "axios";
import { Check, ChevronLeft, X } from "lucide-react";
import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const NotificationPage = () => {
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");
  const [notificationSetting, setNotificationSetting] =
    useState<NotificationSetting | null>(null);
  const { user, getUser, authenticated } = useAuth();
  const navigate = useNavigate();

  const changeSetting = () => {
    if (user == null) return;
    if (notificationSetting == null) return;
    console.log(notificationSetting.notificationSettingId);
    axios
      .post(
        "http://localhost:4000/auth/setting/update",
        {
          userId: user.user_id,
          notificationSettingId: notificationSetting.notificationSettingId,
          emailFollower: notificationSetting.emailFollower,
          emailAlbum: notificationSetting.emailAlbum,
          webFollower: notificationSetting.webFollower,
          webAlbum: notificationSetting.webAlbum,
        },
        {
          withCredentials: true,
        },
      )
      .then((res) => {
        console.log(res);
        setSuccess("Notification setting updated");
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

  useEffect(() => {
    if (user == null) return;
    setNotificationSetting(user.notificationSetting);
  }, [user]);

  return (
    <div className={"wrapper"}>
      {error && <ErrorModal error={error} setError={setError} />}
      {success && <SuccessModal success={success} setSuccess={setSuccess} />}
      <Navbar />
      <div className="container">
        <div className={"adminContainer"}>
          <div className={"adminBox"}>
            <div className={"editProfileTitle"}>
              <Link to={"/account/settings"} className={"link"}>
                <ChevronLeft />
              </Link>
              <h1>Notification settings</h1>
              <p>
                Pick the notification you want to get via push or email. These
                preferences apply to push and email
              </p>
            </div>
            <div className={"verifyContainer"}>
              <div className={"left"}></div>
              <div className={"right"}>
                <div>Email</div>
                <div>Web</div>
              </div>
            </div>
            <div className={"verifyContainer"}>
              <div className={"left"}>
                <div className={"userContent"}>
                  <h6>Followers</h6>
                  <p>Update from new followers</p>
                </div>
              </div>
              <div className={"right"}>
                <div
                // onClick={() => {
                //     onVerify(art.artist.artistId);
                // }}
                >
                  {notificationSetting?.emailFollower ? (
                    <Check
                      className={"check"}
                      onClick={() => {
                        setNotificationSetting((prevState) => {
                          if (prevState === null) return null;
                          return {
                            ...prevState,
                            emailFollower: false,
                          };
                        });
                      }}
                    />
                  ) : (
                    <X
                      className={"x"}
                      onClick={() => {
                        setNotificationSetting((prevState) => {
                          if (prevState === null) return null;
                          return {
                            ...prevState,
                            emailFollower: true,
                          };
                        });
                      }}
                    />
                  )}
                </div>
                <div
                // onClick={() => {
                //     onCancel(art.artist.artistId);
                // }}
                >
                  {notificationSetting?.webFollower ? (
                    <Check
                      className={"check"}
                      onClick={() => {
                        setNotificationSetting((prevState) => {
                          if (prevState === null) return null;
                          return {
                            ...prevState,
                            webFollower: false,
                          };
                        });
                      }}
                    />
                  ) : (
                    <X
                      className={"x"}
                      onClick={() => {
                        setNotificationSetting((prevState) => {
                          if (prevState === null) return null;
                          return {
                            ...prevState,
                            webFollower: true,
                          };
                        });
                      }}
                    />
                  )}
                </div>
              </div>
            </div>
            <div className={"verifyContainer"}>
              <div className={"left"}>
                <div className={"userContent"}>
                  <h6>Album Recommendations</h6>
                  <p>Update from new releases album from artist you follow</p>
                </div>
              </div>
              <div className={"right"}>
                <div
                // onClick={() => {
                //     onVerify(art.artist.artistId);
                // }}
                >
                  {notificationSetting?.emailAlbum ? (
                    <Check
                      className={"check"}
                      onClick={() => {
                        setNotificationSetting((prevState) => {
                          if (prevState === null) return null;
                          return {
                            ...prevState,
                            emailAlbum: false,
                          };
                        });
                      }}
                    />
                  ) : (
                    <X
                      className={"x"}
                      onClick={() => {
                        setNotificationSetting((prevState) => {
                          if (prevState === null) return null;
                          return {
                            ...prevState,
                            emailAlbum: true,
                          };
                        });
                      }}
                    />
                  )}
                </div>
                <div
                // onClick={() => {
                //     onCancel(art.artist.artistId);
                // }}
                >
                  {notificationSetting?.webAlbum ? (
                    <Check
                      className={"check"}
                      onClick={() => {
                        setNotificationSetting((prevState) => {
                          if (prevState === null) return null;
                          return {
                            ...prevState,
                            webAlbum: false,
                          };
                        });
                      }}
                    />
                  ) : (
                    <X
                      className={"x"}
                      onClick={() => {
                        setNotificationSetting((prevState) => {
                          if (prevState === null) return null;
                          return {
                            ...prevState,
                            webAlbum: true,
                          };
                        });
                      }}
                    />
                  )}
                </div>
              </div>
            </div>
            <div className={"saveButton"}>
              <Link to={"/account/settings"}>Cancel</Link>
              <button className={"createMusic"} onClick={changeSetting}>
                Save
              </button>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};
