import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

function Posts() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    fetch(`/api/posts`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setPosts(data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return (
    <div>
      <hr />
      {posts.map((p) => (
        <div>
          <Link to={`/blog/${p.id}`}>{p.title}</Link>
          <p>{p.date}</p>
          <p>{p.summary}</p>
          <hr />
        </div>
      ))}
    </div>
  );
}

export default Posts;
