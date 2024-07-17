import type { AxiosResponse } from "axios";
import axios from "axios";
import { BoxSelect, Camera, Minus, Plus } from "lucide-react";
import type { ChangeEvent } from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import { ControlMusic } from "../component/ControlMusic.tsx";
import { ErrorModal } from "../component/ErrorModal.tsx";
import { Main } from "../component/Main.tsx";
import { RightSideBar } from "../component/RightSideBar.tsx";
import { SideBar } from "../component/SideBar.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const CreateMusicPage = () => {
  const [track, setTrack] = useState<number>(1);
  const { user, authenticated } = useAuth();
  const [success, setSuccess] = useState<string>("");
  const handleAddTrack = () => {
    setTrack((t) => t + 1);
    setTracks((prevTracks) => [
      ...prevTracks,
      { name: "", file: null, duration: 0 },
    ]);
  };

  const handleRemoveTrack = () => {
    if (track === 1) return;
    setTrack((t) => t - 1);
    setTracks((prevTracks) => prevTracks.slice(0, -1));
  };

  const [image, setImage] = useState<File | null>(null);
  const [tracks, setTracks] = useState(
    Array.from(
      { length: track },
      (): { file: File | null; name: string; duration: number } => ({
        name: "",
        file: null,
        duration: 0,
      }),
    ),
  );
  const [title, setTitle] = useState<string>("");
  const [error, setError] = useState<string>("");

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

  const handleInputChange = (
    index: number,
    event: ChangeEvent<HTMLInputElement>,
  ) => {
    const { name, value, files } = event.target;

    const newTracks = [...tracks];
    console.log(index);
    if (name === "track") {
      newTracks[index].name = value;
    } else if (name === "trackFile" && files != null) {
      const file = files[0];
      newTracks[index].file = file;
      if (file.name.endsWith(".mp3")) {
        const reader = new FileReader();

        const currentIndex = index;
        console.log(currentIndex);

        reader.onload = function () {
          const audio = new Audio();
          audio.src = reader.result as string;
          audio.addEventListener("loadedmetadata", function () {
            newTracks[currentIndex].duration = audio.duration;
          });

          audio.addEventListener("error", function () {
            console.error("Error loading audio");
          });
          audio.load();
        };

        reader.readAsDataURL(file);
      }
    }

    setTracks(newTracks);
  };

  const handleInputTitle = (event: ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };

  const navigate = useNavigate();

  const handlePost = () => {
    if (user == null) return;
    let type: string;
    if (track >= 1 && track <= 3) {
      type = "Single";
    } else if (track >= 4 && track <= 6) {
      type = "Eps";
    } else {
      type = "Albums";
    }

    if (title === "") {
      setError("Please fill the title");
      return;
    }

    if (tracks.some((track) => track.name === "" || track.file == null)) {
      setError("Please fill all the track name and upload the file");
      return;
    }

    if (image === null) {
      setError("Please upload the banner");
      return;
    }

    const errs = [];

    axios
      .get("http://localhost:4000/auth/artist/get?id=" + user.user_id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Artist>>) => {
        const artist = res.data.data;
        const formData = new FormData();
        formData.append("image", image as Blob);
        formData.append("title", title);
        formData.append("type", type);
        formData.append("artistId", artist.artistId);
        axios
          .post("http://localhost:4000/artist/album/create", formData, {
            withCredentials: true,
          })
          .then((res: AxiosResponse<WebResponse<Album>>) => {
            console.log(res);
            const albumId = res.data.data.albumId;
            tracks.forEach((track) => {
              if (track.file == null) return;
              const trackData = new FormData();
              trackData.append("artistId", artist.artistId);
              trackData.append("albumId", albumId.toString());
              trackData.append("title", track.name);
              trackData.append("song", track.file);
              trackData.append(
                "duration",
                Math.floor(track.duration).toString(),
              );
              axios
                .post("http://localhost:4000/artist/song/create", trackData, {
                  withCredentials: true,
                })
                .then((res) => {
                  console.log(res);
                  setSuccess("Success Created Album");
                })
                .catch((err: unknown) => {
                  errs.push(err);
                });
            });
          })
          .catch((err: unknown) => {
            console.log(err);
            errs.push(err);
          });
      })
      .catch((err: unknown) => {
        console.log(err);
        errs.push(err);
      });

    if (errs.length > 0) {
      setError("Error creating album");
    }
  };

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

  useEffect(() => {
    if (user == null) return;
    if (user.role != "Artist") navigate("/home");
  }, [user]);

  return (
    <div className={"outer"}>
      {error && <ErrorModal setError={setError} error={error} />}
      {success && <SuccessModal success={success} setSuccess={setSuccess} />}
      <div className={"App"}>
        <SideBar />

        <Main setSearch={null}>
          <div className="profileHeader">
            <div>
              <h1>Create New Music</h1>
            </div>
          </div>
          <div className="newMusicContainer">
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
            <div className={"trackContainer"}>
              <div className={"albumTrackTitle"}>
                <div className={"titleContainer"}>
                  <label htmlFor="title">Title</label>
                  <input
                    type="text"
                    className={"inputText"}
                    id="title"
                    name="title"
                    onChange={handleInputTitle}
                  />
                </div>
                <div className={"collectionContainer"}>
                  <label htmlFor="album">Collection Type</label>
                  <select name="album" id="album" disabled={true}>
                    <option value="Single" selected={track >= 1 && track <= 3}>
                      Singles
                    </option>
                    <option value="Eps" selected={track >= 4 && track <= 6}>
                      Eps
                    </option>
                    <option value="Albums" selected={track > 6}>
                      Albums
                    </option>
                  </select>
                </div>
              </div>
              <h6>Tracks</h6>
              <div className={"trackList"}>
                {Array.from({ length: track }).map((_, index) => (
                  <div className={"track"} key={index}>
                    <div>
                      <label htmlFor="track">#{index + 1}. </label>
                      <input
                        type="text"
                        className={"inputText"}
                        id="track"
                        name="track"
                        placeholder={"Name of track"}
                        onChange={(e) => {
                          handleInputChange(index, e);
                        }}
                      />
                    </div>
                    <label htmlFor={`trackFile${String(index)}`}>
                      <div className={"uploadSong"}>
                        {tracks[index].file ? (
                          <p>Uploaded</p>
                        ) : (
                          <p>Upload MP3 </p>
                        )}
                      </div>
                      <input
                        type="file"
                        id={`trackFile${String(index)}`}
                        className={"trackFile"}
                        name="trackFile"
                        onChange={(e) => {
                          handleInputChange(index, e);
                        }}
                      />
                    </label>
                  </div>
                ))}
              </div>
            </div>
          </div>
          <div className={"plusMinus"}>
            <div className={"logoWrapper"} onClick={handleAddTrack}>
              <Plus />
            </div>
            <div className={"logoWrapper"} onClick={handleRemoveTrack}>
              <Minus />
            </div>
          </div>
          <div className={"saveButton"}>
            <Link to={"/your-post/"}>Cancel</Link>
            <button className={"createMusic"} onClick={handlePost}>
              Post Music
            </button>
          </div>
        </Main>
        <RightSideBar />
      </div>
      <ControlMusic />
    </div>
  );
};
