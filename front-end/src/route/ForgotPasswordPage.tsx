import axios from "axios";
import type { ChangeEvent } from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const ForgotPasswordPage = () => {
  const onChangeInput = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.value != "") {
      setEmail(e.target.value);
    }
  };

  const [email, setEmail] = useState<string>("");
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");
  const searchAccount = () => {
    axios
      .post("http://localhost:4000/user/forgot-password?email=" + email)
      .then((res) => {
        console.log(res);
        setSuccess("Success! Check your email for a reset link");
      })
      .catch((err: unknown) => {
        setError("Email not found");
        console.log(err);
      });
  };

  const { authenticated } = useAuth();
  const navigate = useNavigate();
  useEffect(() => {
    if (authenticated == null) return;
    if (authenticated) {
      navigate("/home");
    }
  }, [authenticated]);

  return (
    <div className={"wrapper"}>
      <Navbar />
      <div className="container">
        {error && <ErrorModal error={error} setError={setError} />}
        {success && <SuccessModal success={success} setSuccess={setSuccess} />}
        <div className={"loginBox"}>
          <h1>Find Your Account</h1>
          <div className="input-group">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              onChange={onChangeInput}
            />
          </div>
          <button className={"loginButton"} onClick={searchAccount}>
            Search
          </button>
          <Link to={"/login"}>Cancel</Link>
        </div>
      </div>
      <Footer />
    </div>
  );
};
