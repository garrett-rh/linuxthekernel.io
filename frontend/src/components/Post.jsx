import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";

const Post = () => {
  const [post, setPost] = useState("");
  let { id } = useParams();

  useEffect(() => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    fetch(`/api/posts/${id}`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setPost(data.content);
      })
      .catch((err) => {
        console.log(err);
      });
  }, [id]);

  return (
    <div>
      <div dangerouslySetInnerHTML={{ __html: post }} />
    </div>
  );
};

export default Post;
