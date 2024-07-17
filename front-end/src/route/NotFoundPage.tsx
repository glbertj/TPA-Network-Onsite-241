import { Footer } from "../component/Footer.tsx";
import { Navbar } from "../component/Navbar.tsx";

export const NotFoundPage = () => {
  return (
    <div className={"wrapper"}>
      <Navbar />
      <div className="container">
        <div className={"adminContainer"}>
          <div className={"adminBox"}>
            <div className={"editProfileTitle"}>
              <h1>404 Page Not Found</h1>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};
