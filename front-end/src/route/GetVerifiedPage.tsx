import type { AxiosError } from "axios";
import axios, { type AxiosResponse } from "axios";
import { BoxSelect, Camera, ChevronLeft } from "lucide-react";
import { type ChangeEvent, useEffect, useState } from "react";
import * as React from "react";
import { Link, useNavigate } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const GetVerifiedPage = () => {
  const { user, authenticated } = useAuth();
  const navigate = useNavigate();
  const [description, setDescription] = useState<string>("");
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");
  const [image, setImage] = useState<File | null>(null);

  const onChangeInput = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setDescription(e.target.value);
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

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

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
    formData.append("userId", user.user_id);
    axios
      .post("http://localhost:4000/auth/artist/create", formData, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Album>>) => {
        console.log(res);
        setSuccess("Request is Formed");
      })
      .catch((err: unknown) => {
        const error = err as AxiosError<WebResponse<string>>;
        if (error.response == undefined) return;
        setError(error.response.data.message);
        console.log(err);
      });
  };

  console.log(user);

  return (
    <div className={"wrapper"}>
      {error && <ErrorModal error={error} setError={setError} />}
      {success && <SuccessModal setSuccess={setSuccess} success={success} />}
      <Navbar />
      <div className="container">
        <div className={"loginBox"}>
          <div className={"editProfileTitle"}>
            <Link to={"/account/settings"}>
              <ChevronLeft />
            </Link>
            <h1>Get Verified</h1>
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
                      Upload Banner Image
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
                <p>Current Role</p>
                <h6>{user?.role}</h6>
              </div>
              <div className={"areaAbout"}>
                <label htmlFor="about">About You</label>
                <textarea id="about" name="about" onChange={onChangeInput} />
              </div>
              <div className={"saveButton"}>
                <Link to={"/account/settings"}>Cancel</Link>
                <button className={"loginButton"} onClick={handlePost}>
                  Get Verified
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
