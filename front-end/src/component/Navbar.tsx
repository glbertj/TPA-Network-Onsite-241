import { ChevronDown, ChevronUp } from "lucide-react";
import { useState } from "react";
import { Link } from "react-router-dom";

import { useAuth } from "../context/UseAuth.tsx";

export const Navbar = () => {
  const { user, logout } = useAuth();
  const [isDrop, setIsDrop] = useState(false);

  return (
    <nav>
      <div className={"left"}>
        <Link to={"/home"}>
          <img className={"logo"} src={"/assets/NJOTIFY.png"} alt={""} />
        </Link>
      </div>
      <div className={"right"}>
        <div className="dropdown">
          <div
            className={"profile"}
            onClick={() => {
              setIsDrop(!isDrop);
            }}
          >
            {user && (
              <img
                src={user.avatar ? user.avatar : "/assets/download (6).png"}
                alt={"p"}
              ></img>
            )}
            <div></div>
            <p>Profile</p>
            {!isDrop ? <ChevronDown /> : <ChevronUp />}
          </div>
          <div className={`dropdown-content ${isDrop ? "active" : ""}`}>
            {user && (
              <>
                <Link to={"/profile/" + user.user_id} className={"link"}>
                  Profile
                </Link>
                <hr className={"hr"} />
                <Link to={"#"} className={"link"} onClick={logout}>
                  Sign Out
                </Link>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
};
