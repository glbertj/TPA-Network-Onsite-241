import { useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";

import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const SettingPage = () => {
  const { logout, authenticated } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

  return (
    <div className={"wrapper"}>
      <Navbar />
      <div className="container">
        <div className={"loginBox"}>
          <div className={"accountLink"}>
            <h1>Account</h1>
            <div className={"account"}>
              <div className={"linkBackground"}>
                <Link to={"/user/edit"}>Edit Profile</Link>
              </div>
              <div className={"linkBackground"}>
                <Link to={"/get-verified"}>Get Verified</Link>
              </div>
            </div>
            <div className={"account"}>
              <div className={"linkBackground"}>
                <Link to={"/notification/setting"}>Notification Settings</Link>
              </div>
              <div className={"linkBackground"}>
                <p style={{ cursor: "pointer" }} onClick={logout}>
                  Sign out
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};
