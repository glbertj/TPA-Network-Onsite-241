import type { AxiosResponse } from "axios";
import axios from "axios";
import { ChevronLeft, ChevronRight, MicVocal, Search } from "lucide-react";
import type { ChangeEvent, Dispatch, SetStateAction } from "react";
import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import { useAuth } from "../context/UseAuth.tsx";
import { ErrorModal } from "./ErrorModal.tsx";

interface IProps {
  result: string;
}

export const Header = ({
  setSearch,
  setIsLoad,
  search,
}: {
  setSearch: Dispatch<SetStateAction<string>> | null;
  setIsLoad?: Dispatch<SetStateAction<boolean>>;
  search?: string;
}) => {
  const [isDrop, setIsDrop] = useState(false);
  const [history, setHistory] = useState<number>(window.history.length);
  const { user, logout } = useAuth();
  useEffect(() => {
    window.onpopstate = () => {
      setHistory(window.history.length);
    };
  }, [window.history.length]);

  const handleBack = () => {
    console.log(window.history.length);
    if (history < 2) return;
    window.history.back();
    setHistory(window.history.length);
  };

  const handleForward = () => {
    if (
      window.history.state !== null &&
      window.history.state !== undefined &&
      history > window.history.state.index
    ) {
      window.history.forward();
    } else {
      console.log("No forward history");
    }
    setHistory(window.history.length);
  };

  const handleSearch = (e: ChangeEvent<HTMLInputElement>) => {
    if (setSearch === null) return;
    setSearch(e.target.value);
  };

  const [error, setError] = useState<string>("");
  const navigate = useNavigate();

  const handleSubmit = (e: ChangeEvent<HTMLInputElement>) => {
    const music = e.target.files?.[0];
    if (music === undefined) {
      setError("Please select a file");
      return;
    }

    if (!music.type.includes("audio")) {
      setError("Please select a mp3 file");
      return;
    }
    if (setIsLoad === undefined) return;
    setIsLoad(true);
    const dataForm = new FormData();
    dataForm.append("files", music);
    axios
      .post(`http://127.0.0.1:5000/getresult`, dataForm, {
        headers: {
          "Content-Type": "multipart/form-data",
          "Access-Control-Allow-Origin": "*",
        },
      })
      .then((res: AxiosResponse<IProps>) => {
        if (setSearch === null) return;
        // console.log(res.data.result);
        setSearch(res.data.result);
        navigate("/search?query=" + res.data.result);
      })
      .catch((err: unknown) => {
        console.log(err);
        setError(err as string);
        setIsLoad(false);
      });
  };

  return (
    <header>
      {error && <ErrorModal error={error} setError={setError} />}
      <div className={"left"}>
        <ChevronLeft
          className={history > 2 ? "disabled" : ""}
          onClick={handleBack}
        />
        <ChevronRight
          className={
            window.history.state !== null &&
            window.history.state !== undefined &&
            history > window.history.state.index + 1
              ? ""
              : "disabled"
          }
          onClick={handleForward}
        />
        {window.location.pathname === "/search" && (
          <div className="search-container">
            <input
              type="text"
              placeholder="Search..."
              onChange={handleSearch}
              value={search}
            />
            <Search />
          </div>
        )}
        {setIsLoad && (
          <div>
            <div>
              <label htmlFor={"file"}>
                <MicVocal />
              </label>
              <input
                type="file"
                id={"file"}
                onChange={handleSubmit}
                style={{ opacity: 0 }}
              />
            </div>
          </div>
        )}
      </div>
      <div className={"right"}>
        <div className="dropdown">
          <img
            src={user?.avatar ?? "/assets/download (6).png"}
            alt={"p"}
            className="profile"
            onClick={() => {
              setIsDrop(!isDrop);
            }}
          ></img>
          <div className={`dropdown-content ${isDrop ? "active" : ""}`}>
            {user && (
              <Link to={"/profile/" + user.user_id} className={"link"}>
                Profile
              </Link>
            )}
            <Link to={"/account/settings"} target={"_blank"} className={"link"}>
              Manage Account
            </Link>
            {user?.role === "Admin" && (
              <Link to={"/artist/verif/"} className={"link"}>
                Verify Artist
              </Link>
            )}
            <hr className={"hr"} />
            <p className={"link"} onClick={logout}>
              Logout
            </p>
          </div>
        </div>
      </div>
    </header>
  );
};
