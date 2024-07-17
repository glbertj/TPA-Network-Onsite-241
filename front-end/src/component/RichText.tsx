import type { AxiosResponse } from "axios";
import axios from "axios";
import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export const RichText = ({ description }: { description: string }) => {
  const [users, setUsers] = useState<User[]>();
  const navigate = useNavigate();
  useEffect(() => {
    axios
      .get("http://localhost:4000/auth/user/get-all", {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<User[]>>) => {
        setUsers(res.data.data);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  }, []);

  const handleTag = (user: User) => {
    if (user.role == "Artist") {
      navigate("/artist/" + user.user_id);
    } else if (user.role == "Listener") {
      navigate("/profile/" + user.user_id);
    } else if (user.role == "Admin") {
      navigate("/profile/" + user.user_id);
    }
  };

  const handleHashTag = (tag: string) => {
    navigate("/search?query=" + tag);
  };

  const convertToRichText = () => {
    const elements = [];
    let curr = 0;

    let texts = "";
    while (curr < description.length) {
      if (description.slice(curr, curr + 8) === "https://") {
        if (texts) {
          elements.push(<span key={curr}>{texts}</span>);
          texts = "";
        }
        const start = curr;
        while (curr < description.length && /\S/.test(description[curr])) {
          curr++;
        }
        const link = description.slice(start, curr);
        elements.push(
          <Link key={curr} to={link} style={{ color: "blue" }} target="_blank">
            {link}
          </Link>,
        );
      } else if (description[curr] === "@") {
        const user = users?.find(
          (user) =>
            user.username ==
            description.slice(curr + 1, curr + 1 + user.username.length),
        );
        console.log(description.slice(curr + 1, curr + 1 + 2));
        if (user != null) {
          if (texts) {
            elements.push(<span key={curr}>{texts}</span>);
            texts = "";
          }
          elements.push(
            <span
              key={curr}
              style={{ color: "green", cursor: "pointer" }}
              onClick={() => {
                handleTag(user);
              }}
            >
              @{user.username}
            </span>,
          );
          curr += user.username.length + 1;
        } else {
          texts += description[curr];
          curr++;
        }
      } else if (description[curr] === "#") {
        if (texts) {
          elements.push(<span key={curr}>{texts}</span>);
          texts = "";
        }
        const start = curr;
        curr++;
        while (curr < description.length && /\w/.test(description[curr])) {
          curr++;
        }
        const text = description.slice(start, curr);
        const tag = text.slice(1);
        elements.push(
          <span
            key={curr}
            style={{ color: "orange", cursor: "pointer" }}
            onClick={() => {
              handleHashTag(tag);
            }}
          >
            {text}
          </span>,
        );
      } else {
        texts += description[curr];
        curr++;
      }
    }

    if (texts) {
      elements.push(<span key={curr}>{texts}</span>);
    }

    return elements;
  };

  return <div>{users && convertToRichText()}</div>;
};
