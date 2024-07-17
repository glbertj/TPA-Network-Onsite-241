import { Music } from "lucide-react";

export const InputTrack = () => {
  return (
    <div className={"track"}>
      <div>
        <label htmlFor="track">#1. </label>
        <input
          type="text"
          className={"inputText"}
          id="track"
          name="track"
          placeholder={"Name of track"}
        />
      </div>
      <label htmlFor="trackFile">
        <div className={"uploadSong"}>
          <p>Upload MP3</p>
          <Music />
        </div>
        <input type="file" id="trackFile" name="trackFile" />
      </label>
    </div>
  );
};
