import type { ChangeEvent } from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const LoginPage = () => {
  const { login, googleLogin, authenticated, error, setError } = useAuth();
  const [inputLogin, setInputLogin] = useState<LoginProps>({} as LoginProps);
  const navigate = useNavigate();

  const onLogin = () => {
    login(inputLogin);
  };

  const onChangeInput = (e: ChangeEvent<HTMLInputElement>) => {
    setInputLogin({ ...inputLogin, [e.target.name]: e.target.value });
  };

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
        <div className={"loginBox"}>
          <h1>Login</h1>
          <button
            className={"google"}
            onClick={() => {
              googleLogin();
            }}
          >
            Continue with Google
          </button>
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
          <button className={"loginButton"} onClick={onLogin}>
            Login
          </button>
          <Link to={"/forgot"}>Forgot your Password?</Link>
          <hr />
          <p>
            Dont have an account?{" "}
            <Link to={"/register"}>Sign Up For NJ Notify</Link>
          </p>
        </div>
      </div>
      <Footer />
    </div>
  );
};
