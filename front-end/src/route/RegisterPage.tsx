import { useEffect, useState } from "react";
import * as React from "react";
import { Link, useNavigate } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const RegisterPage = () => {
  const { register, googleLogin, success, setSuccess, error, setError } =
    useAuth();
  // const {theme,setTheme} = useTheme()
  const [registerLogin, setRegisterLogin] = useState<RegisterProps>({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  } as RegisterProps);
  const [err, setErr] = useState<string>("");
  const onRegister = () => {
    if (
      registerLogin.username == "" ||
      registerLogin.email == "" ||
      registerLogin.password == "" ||
      registerLogin.confirmPassword == ""
    ) {
      setError("Please fill in all fields");
      return;
    }

    if (registerLogin.password != registerLogin.confirmPassword) {
      setError("Passwords do not match");
      return;
    }

    if (registerLogin.password.length < 8) {
      setError("Password must be at least 8 characters");
      return;
    }

    const lower = registerLogin.password.match(/[a-z]/);
    const upper = registerLogin.password.match(/[A-Z]/);
    const numeric = registerLogin.password.match(/[0-9]/);
    if (!lower || !upper || !numeric) {
      setError(
        "Password must contain at least one uppercase, lowercase, and number",
      );
      return;
    }

    register(registerLogin);
  };

  const onChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setRegisterLogin({ ...registerLogin, [e.target.name]: e.target.value });
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
      {error && <ErrorModal error={error} setError={setError} />}
      {err && <ErrorModal error={err} setError={setErr} />}
      {success && <SuccessModal success={success} setSuccess={setSuccess} />}
      <Navbar />
      <div className="container">
        <div className={"loginBox"}>
          <h1>Sign up to start Listening</h1>
          <button
            className={"google"}
            onClick={() => {
              googleLogin();
            }}
          >
            Continue with Google
          </button>
          <div className="input-group">
            <label htmlFor="username">Username</label>
            <input
              type="text"
              id="username"
              name="username"
              onChange={onChangeInput}
            />
          </div>
          <div className="input-group">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              onChange={onChangeInput}
            />
          </div>
          <div className="input-group">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              onChange={onChangeInput}
            />
          </div>
          <div className="input-group">
            <label htmlFor="password">Confirm Password</label>
            <input
              type="password"
              id="confirmPassword"
              name="confirmPassword"
              onChange={onChangeInput}
            />
          </div>
          <button className={"loginButton"} onClick={onRegister}>
            Sign Up
          </button>
          <hr />
          <p>
            Already have an account?{" "}
            <Link to={"/login"}>Log in to NJ Notify</Link>
          </p>
        </div>
      </div>
      <Footer />
    </div>
  );
};
