import { Link } from "react-router-dom";

function About() {
  return (
    <>
      <p className="fs-1 text-center">About Me</p>
      <hr />
      <div className="row">
          <div className="col-md-6">
              <img src="https://d17x1veniq4brh.cloudfront.net/fishing.jpg" className="img-fluid img-thumbnail p-3"
                   alt="Fishing"/>
          </div>
          <div className="col-md-6">
              <p>
                  I'm Garrett Harvey. I am a DevOps & Software Engineer. This is a
            website that I made for fun. I intend to make this mostly about my
            technical work and personal projects but more than likely it'll just
            end up being blog posts about my time spent fishing.
          </p>
          <Link to="/blog/AboutMe">For more information on who I am and how I started, check this post.</Link>
          <p>You can also find my LinkedIn and GitHub linked in the top right! My most up to date resume can be found on my LinkedIn profile.</p>
        </div>
      </div>
    </>
  );
}

export default About;
