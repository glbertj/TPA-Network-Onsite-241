import axios, { type AxiosError, type AxiosResponse } from "axios";
import { useEffect, useState } from "react";
import * as React from "react";
import { Link, useNavigate, useSearchParams } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const ResetPasswordPage = () => {
  const [searchParams] = useSearchParams();
  const id = searchParams.get("id");
  const navigate = useNavigate();
  const [user, setUser] = useState<User | null>(null);
  const [resetPass, setResetPass] = useState<ResetPass>({
    confirmPassword: "",
    password: "",
  } as ResetPass);
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");

  useEffect(() => {
    if (id != null) {
      axios
        .get("http://localhost:4000/user/valid-verify?id=" + id, {
          headers: {
            "Content-Type": "application/json",
          },
        })
        .then((res: AxiosResponse<WebResponse<User>>) => {
          console.log(res);
          setUser(res.data.data);
        })
        .catch((err: unknown) => {
          console.log(err);
          navigate("/login");
        });
    }
  }, [id]);

  const updatePassword = () => {
    if (user == null) return;
    if (resetPass.password == "" || resetPass.confirmPassword == "") {
      setError("Please fill in all fields");
      return;
    }

    if (resetPass.password != resetPass.confirmPassword) {
      setError("Passwords do not match");
      return;
    }

    if (resetPass.password.length < 8) {
      setError("Password must be at least 8 characters");
      return;
    }

    const lower = resetPass.password.match(/[a-z]/);
    const upper = resetPass.password.match(/[A-Z]/);
    const numeric = resetPass.password.match(/[0-9]/);
    if (!lower || !upper || !numeric) {
      setError(
        "Password must contain at least one uppercase, lowercase, and number",
      );
      return;
    }

    axios
      .post("http://localhost:4000/user/reset-password", {
        userId: user.user_id,
        password: resetPass.password,
      })
      .then((res) => {
        console.log(res);
        navigate("/login");
      })
      .catch((err: unknown) => {
        const error = err as AxiosError<WebResponse<string>>;
        if (error.response == undefined) return;
        setError(error.response.data.message);
      });
  };

  const onChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setResetPass({ ...resetPass, [e.target.name]: e.target.value });
  };

  const { authenticated } = useAuth();
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
        {success && <SuccessModal success={success} setSuccess={setSuccess} />}
        {error && <ErrorModal error={error} setError={setError} />}
        <div className={"loginBox"}>
          <h1>Reset Password</h1>
          <div className="input-group">
            <label htmlFor="password">New Password</label>
            <input
              type="password"
              id="password"
              name="password"
              onChange={onChangeInput}
            />
          </div>
          <div className="input-group">
            <label htmlFor="confirmPassword">Confirm Password</label>
            <input
              type="password"
              id="confirmPassword"
              name="confirmPassword"
              onChange={onChangeInput}
            />
          </div>
          <button className={"loginButton"} onClick={updatePassword}>
            Reset Password
          </button>
          <Link to={"/login"}>Cancel</Link>
        </div>
      </div>
      <Footer />
    </div>
  );
};
