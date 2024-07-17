import type { AxiosError } from "axios";
import axios, { type AxiosResponse } from "axios";
import { BoxSelect, Camera, ChevronLeft } from "lucide-react";
import { type ChangeEvent, useEffect, useState } from "react";
import * as React from "react";
import { Link, useNavigate } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { RichText } from "../component/RichText.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { useSong } from "../context/UseSong.tsx";

export const CreatePlaylistPage = () => {
  const { user } = useAuth();
  const { updatePlaylist } = useSong();

  const [description, setDescription] = useState<string>("");
  const [title, setTitle] = useState<string>("");
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");
  const [image, setImage] = useState<File | null>(null);

  const onChangeDesc = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setDescription(e.target.value);
  };

  const onChangeTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };

  const handleChangeImage = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files == null) return;
    if (e.target.files.length === 0) {
      setError("Please select an image");
      return;
    } else if (
      !e.target.files[0].name.endsWith(".jpg") &&
      !e.target.files[0].name.endsWith(".png") &&
      !e.target.files[0].name.endsWith(".jpeg")
    ) {
      setError("Please select an image with jpg, jpeg, or png format");
      return;
    }
    setImage(e.target.files[0]);
  };

  const handlePost = () => {
    if (user == null) return;

    if (description === "") {
      setError("Please fill the Description field");
      return;
    }

    if (image === null) {
      setError("Please upload the banner");
      return;
    }

    const formData = new FormData();
    formData.append("image", image as Blob);
    formData.append("description", description);
    formData.append("title", title);
    formData.append("userId", user.user_id);

    axios
      .post("http://localhost:4000/auth/playlist/create", formData, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Album>>) => {
        console.log(res);
        updatePlaylist();
        setSuccess("Playlist created");
      })
      .catch((err: unknown) => {
        const error = err as AxiosError<WebResponse<string>>;
        if (error.response == undefined) return;
        setError(error.response.data.message);
      });
  };

  const navigate = useNavigate();
  useEffect(() => {
    if (user == null) return;
  }, [user]);

  return (
    <div className={"wrapper"}>
      {error && <ErrorModal error={error} setError={setError} />}
      {success && <SuccessModal success={success} setSuccess={setSuccess} />}
      <Navbar />
      <div className="container">
        <div className={"loginBox"}>
          <div className={"editProfileTitle"}>
            <ChevronLeft
              onClick={() => {
                navigate("/home");
              }}
            />
            <h1>Create Playlist</h1>
          </div>
          <div className={"inputVerify"}>
            <label htmlFor="image">
              <div className={"uploadImage"}>
                {image ? (
                  <>
                    <img
                      src={URL.createObjectURL(image)}
                      alt={"banner"}
                      style={{ width: "200px", height: "200px" }}
                    />
                    <input
                      type={"file"}
                      id={"image"}
                      onChange={handleChangeImage}
                    />
                  </>
                ) : (
                  <>
                    <div className={"camera"}>
                      <Camera />
                      Upload Playlist Image
                      <input
                        type={"file"}
                        id={"image"}
                        onChange={handleChangeImage}
                      />
                    </div>
                    <BoxSelect className={"boxSelect"} />
                  </>
                )}
              </div>
            </label>
            <div className={"verify"}>
              <div className={"role"}>
                <label htmlFor="title">Playlist Title</label>
                <input
                  type={"text"}
                  name={"title"}
                  id={"title"}
                  onChange={onChangeTitle}
                />
              </div>
              <div className={"areaAbout"}>
                <label htmlFor="about">Playlist Description</label>
                <textarea
                  id="about"
                  name="about"
                  onChange={onChangeDesc}
                  style={{
                    border: "1px solid #ccc",
                    padding: "10px",
                    minHeight: "100px",
                    overflow: "auto",
                  }}
                />
                <RichText description={description} />
              </div>
              <div className={"saveButton"}>
                <Link to={"/home"}>Cancel</Link>
                <button className={"loginButton"} onClick={handlePost}>
                  Create Playlist
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};
