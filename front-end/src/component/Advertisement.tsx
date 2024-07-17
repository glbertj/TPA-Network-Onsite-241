import { BadgeCheck, ChevronRight } from "lucide-react";
import { Link } from "react-router-dom";

import { useSong } from "../context/UseSong.tsx";

export const Advertisement = () => {
  const { advertise } = useSong();
  return (
    <>
      <div className="rightSideBarHeader">
        <h4>Your music will continue after the break</h4>
        {/*<X*/}
        {/*  onClick={() => {*/}
        {/*    showDetailHandler("");*/}
        {/*  }}*/}
        {/*/>*/}
      </div>
      <div className="trackImage">
        <img src={advertise?.image} alt="Song Cover" />
      </div>
      <div className="songTitle">
        <div>
          <h3>{advertise?.publisherName}</h3>
          <p>Advertisement</p>
        </div>
        <div>
          <BadgeCheck />
        </div>
      </div>
      <Link to={"https://ads.spotify.com/en-US/"}>
        <div className={"learnMore"}>
          <h3>Learn More</h3>
          <ChevronRight />
        </div>
      </Link>
    </>
  );
};
