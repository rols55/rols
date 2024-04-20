"use client";
import { useState } from "react";
import { useRouter} from "next/navigation"

const initialValues = {
  text: "",
  images: [],
};

export default function Reply({postId}) {

  const router = useRouter();

  const [formData, setFormData] = useState({
    text: "",
  });

  const [imageFile, setImageFile] = useState(null);

  const handleImageChange = (event) => {
    const file = event.target.files[0];
    setImageFile(file);
  };

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: value,
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
  
    const payload = new FormData();
    payload.append("postid", postId);
    //console.log(postId);
  
    for (const key in formData) {
        payload.append(key, formData[key]);
    }
  
    if (imageFile) {
        payload.append("image", imageFile);
    }
  
    const options = {
        method: "POST",
        credentials: 'include',
        body: payload,
    };
  
    try {
        const response = await fetch(
            `http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/comment/create`,
            options
        );
        if (response.ok) {
            const data = await response.json();
            window.location.reload();

        } else {
            const statusMsg = await response.text();
            console.log(statusMsg);
        }
    } catch (error) {
        console.error(error);
    }
  }

  return (
    <form
      className="reply"
      onSubmit={handleSubmit}
      style={{ display: "flex", flexDirection: "column", margin: 20, gap: 5 }}
    >
      <label htmlFor="text">Text</label>
      <textarea
        style={{ height: 100 }}
        id="text"
        name="text"
        type="textarea"
        value={formData.text}
        onChange={handleChange} required
      />

      <label htmlFor="images">Add picture (optional)</label>
      <input
        className="button-27"
        style={{ maxWidth: "300px" }}
        id="imageInput"
        name="image"
        type="file"
        onChange={handleImageChange}
      />

      <button type="submit">Submit</button>
    </form>
  );
};
