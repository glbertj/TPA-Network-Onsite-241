import { Instagram, Linkedin, Twitter } from "lucide-react";
import { useNavigate } from "react-router-dom";

export const Footer = () => {
  const navigate = useNavigate();

  const handleAbout = () => {
    navigate("https://www.spotify.com/id-id/about-us/contact/");
  };
  const handleJob = () => {
    navigate("https://www.lifeatspotify.com/");
  };

  const handleRecord = () => {
    navigate("https://newsroom.spotify.com/");
  };

  const handleArtist = () => {
    navigate("https://artists.spotify.com/home");
  };

  const handleInstagram = () => {
    navigate("https://www.instagram.com/spotify/");
  };

  return (
    <footer>
      <div className={"up"}>
        <div className={"left"}>
          <img src={"/assets/NJOTIFY.png"} alt={""} />
          <div className={"footerContent"}>
            <p className={"footerTitle"}>COMPANY</p>
            <p onClick={handleAbout}>About</p>
            <p onClick={handleJob}>Work</p>
            <p onClick={handleRecord}>For the Record</p>
          </div>
          <div className={"footerContent"}>
            <p className={"footerTitle"}>COMMUNITY</p>
            <p onClick={handleArtist}>For Artist</p>
            <p>Developer</p>
            <p>Advertisement</p>
            <p>Investor</p>
            <p>Vendor</p>
          </div>

          <div className={"footerContent"}>
            <p className={"footerTitle"}>NJOTIFY PACKAGES</p>
            <p>Individual</p>
            <p>Premium Duo</p>
            <p>Premium Student</p>
          </div>
        </div>
        <div className={"right"}>
          <div>
            <Twitter />
          </div>
          <div onClick={handleInstagram}>
            <Instagram />
          </div>
          <div>
            <Linkedin />
          </div>
        </div>
      </div>
      <div className={"down"}>
        <div className={"left"}>
          <p>Privacy Policy</p>
          <p>Terms of Service</p>
        </div>
        <div className={"right"}>
          <p>© 2024 NJ NOTIFY</p>
        </div>
      </div>
    </footer>
  );
};
