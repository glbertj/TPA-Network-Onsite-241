import { useGoogleLogin } from "@react-oauth/google";
import type { AxiosResponse } from "axios";
import axios,{ type AxiosError } from "axios";
import { useAtom } from "jotai";
import { atomWithStorage } from "jotai/utils";
import type { Dispatch, ReactNode, SetStateAction } from "react";
import { createContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const users = atomWithStorage<User | null>("user", null);

interface IProps {
  user: User | null;
  login: (user: LoginProps) => void;
  logout: () => void;
  register: (user: RegisterProps) => void;
  error: string;
  googleLogin: () => void;
  getUser: () => void;
  authenticated: boolean | null;
  success: string;
  setSuccess: Dispatch<SetStateAction<string>>;
  setError: Dispatch<SetStateAction<string>>;
}

export const AuthContext = createContext<IProps>({} as IProps);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [user, setUser] = useAtom(users);
  const [error, setError] = useState<string>("");
  const navigate = useNavigate();
  const [authenticated, setAuthenticated] = useState<boolean | null>(null);
  const [success, setSuccess] = useState<string>("");

  useEffect(() => {
    getUser();
    setError("");
  }, []);

  const getUser = () => {
    axios
      .get("http://localhost:4000/auth/user/current-user", {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<User>>) => {
        setUser(res.data.data);
        setAuthenticated(true);
        const userId = res.data.data.user_id;
        const eventSource = new EventSource(
          "http://localhost:4000/auth/sse/notification-stream?id=" + userId,
          { withCredentials: true },
        );

        eventSource.onopen = function () {
          console.log("Connection to server opened.");
        };

        eventSource.addEventListener(
          "notif-updated",
          function (event: MessageEvent<string>) {
            // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
            const data: Notification = JSON.parse(event.data);
            new Notification(data.title, {
              body: data.body,
            });
          },
        );

        eventSource.addEventListener("initial", function (event) {
          console.log("Received initial message:", event.data);
        });

        eventSource.onerror = function (error) {
          console.error("EventSource error:", error);
          eventSource.close();
          setTimeout(() => {
            getUser();
          }, 5000);
        };

        if ("Notification" in window) {
          Notification.requestPermission()
            .then((permission) => {
              if (permission === "granted") {
                console.log("Notification permission granted.");
              } else {
                console.log("Notification permission denied.");
              }
            })
            .catch((error: unknown) => {
              console.error("Error requesting notification permission:", error);
            });
        } else {
          console.log("This browser does not support notifications.");
        }
      })
      .catch((error: unknown) => {
        setAuthenticated(false);
        console.error("Error fetching user:", error);
      });
  };

  const login = (user: LoginProps) => {
    axios
      .post(
        "http://localhost:4000/user/login",
        {
          Email: user.email,
          Password: user.password,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true,
        },
      )
      .then((response) => {
        console.log(response.data);
        setAuthenticated(true);
        getUser();
        navigate("/home");
      })
      .catch((error: unknown) => {
        setError("Invalid email or password");
        console.error("Error fetching or processing data:", error);
      });
  };

  const register = (user: RegisterProps) => {
    axios
      .put("http://localhost:4000/user/register", {
        Email: user.email,
        Password: user.password,
        Username: user.username,
      })
      .then((response) => {
        console.log(response.data);
        setSuccess("User registered successfully! Please Check your Email");
      })
      .catch((err: unknown) => {
        const error = err as AxiosError<WebResponse<string>>;
        if (error.response == undefined) return;
        setError(error.response.data.message);
        console.error("There was an error registering the user!", err);
      });
  };

  const googleLogin = useGoogleLogin({
    onSuccess: (codeResponse) => {
      axios
        .get(
          `http://localhost:4000/auth/google/callback?code=${codeResponse.code}`,
          {
            headers: {
              "Content-Type": "application/json",
            },
            withCredentials: true,
          },
        )
        .then((res: AxiosResponse<WebResponse<User>>) => {
          console.log(res.data.data);
          getUser();
          navigate("/home");
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    },
    flow: "auth-code",
  });

  const logout = () => {
    if (user == null) return;
    void axios
      .get("http://localhost:4000/auth/user/logout?id=" + user.user_id, {
        withCredentials: true,
      })
      .then(() => {
        setUser(null);
        setAuthenticated(false);
        navigate("/login");
      });
  };

  const values: IProps = {
    user,
    error,
    login,
    register,
    logout,
    googleLogin,
    getUser,
    authenticated,
    success,
    setSuccess,
    setError,
  };
  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>;
};
