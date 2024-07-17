import axios from "axios";
import { ChevronLeft } from "lucide-react";
import { useEffect, useState } from "react";
import * as React from "react";
import { Link, useNavigate } from "react-router-dom";

import { ErrorModal } from "../component/ErrorModal.tsx";
import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";
import { SuccessModal } from "../component/SuccessModal.tsx";
import { useAuth } from "../context/UseAuth.tsx";

export const EditProfilePage = () => {
  const { user, authenticated } = useAuth();
  const navigate = useNavigate();

  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");
  const [editProps, setEditProps] = useState<EditProps>({} as EditProps);

  const onChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEditProps({ ...editProps, [e.target.name]: e.target.value });
    console.log(editProps);
  };

  const onChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setEditProps({ ...editProps, [e.target.name]: e.target.value });
    console.log(editProps);
  };

  const onEdit = () => {
    if (user == null) return;
    void axios
      .post(
        "http://localhost:4000/auth/user/edit-prof",
        {
          userId: user.user_id,
          country: editProps.country,
          dob: new Date(editProps.dob),
          gender: editProps.gender,
        },
        {
          withCredentials: true,
        },
      )
      .then((res) => {
        setSuccess("Profile Updated");
        console.log(res);
      });
  };

  useEffect(() => {
    if (user == null) return;
    setEditProps({
      ...editProps,
      userId: user.user_id,
      dob: user.dob,
      country: user.country,
      gender: user.gender,
    });
  }, [user]);

  useEffect(() => {
    if (authenticated == null) return;
    if (!authenticated) {
      navigate("/login");
    }
  }, [authenticated]);

  return (
    <div className={"wrapper"}>
      {error && <ErrorModal error={error} setError={setError} />}
      {success && <SuccessModal setSuccess={setSuccess} success={success} />}
      <Navbar />
      <div className="container">
        <div className={"loginBox"}>
          <div className={"editProfileTitle"}>
            <Link to={"/account/settings"}>
              <ChevronLeft />
            </Link>
            <h1>Edit Profile</h1>
            <p>User ID</p>
            <p>{user?.user_id}</p>
          </div>
          <div className="input-group">
            <label htmlFor="email">Email</label>
            <input
              disabled={true}
              value={user?.email}
              type="email"
              id="email"
              name="email"
              style={{ cursor: "not-allowed", opacity: 0.7 }}
            />
          </div>
          <div className="input-group">
            <label htmlFor="gender">Gender</label>
            <select
              value={editProps.gender}
              id="gender"
              name="gender"
              onChange={onChangeSelect}
            >
              <option value="">Select Gender</option>
              <option value="Male">Male</option>
              <option value="Female">Female</option>
            </select>
          </div>
          <div className={"dobInput"}>
            <div className="input-group">
              <label htmlFor="dob">Date of Birth</label>
              <input
                type="date"
                value={
                  editProps.dob &&
                  new Date(editProps.dob).toISOString().split("T")[0]
                }
                id="dob"
                name="dob"
                onChange={onChangeInput}
              />
            </div>
            <div className="input-group">
              <label htmlFor="country">Country</label>
              <input
                type="text"
                value={editProps.country}
                id="country"
                name="country"
                onChange={onChangeInput}
              />
            </div>
          </div>
          <div className={"saveButton"}>
            <Link to={"/account/settings"}>Cancel</Link>
            <button className={"loginButton"} onClick={onEdit}>
              Save Profile
            </button>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};
