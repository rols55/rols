"use client";
//import { useState } from "react";
import { useRouter} from "next/navigation"
import React, { useState, useEffect } from 'react';


const initialValues = {
  title: "",
  text: "",
  images: [],
};

let followersArray = [""];

export default function Form ({groupId}) {

  const router = useRouter();

  const [formData, setFormData] = useState({
    title: "",
    text: "",
    privacy: "",
    followers: ""
  });

  const [formErrors, setFormErrors] = useState({
    title: "",
    text: "",
    privacy: "",
    followers: ""
  });

  const [followersArray, setFollowersArray] = useState([""]); 

  useEffect(() => {
    const fetchFollowers = async () => {
      try {
        const response = await fetch(`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/followers`, {
          credentials: 'include'
        });
        if (response.ok) {
          const data = await response.json();
          const followers = data.followers
            .filter((follower) => follower.accepted === true)
            .map((follower) => follower.username);
          setFollowersArray(followers); 
        } else {
          const statusMsg = await response.text();
          console.log(statusMsg);
        }
      } catch (error) {
        console.error(error);
      }
    };

    fetchFollowers();
  }, []); 
   
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

    /*const err = getErr(name, value);
    setFormErrors((prevErrs) => ({
        ...prevErrs,
        [name]: err,
    }));*/
};
/*
const getErr = (name, value) => {
  switch (name) {
      case "title":
          if (!value.trim()) {
              return "This field is required";
          }
          break;
      case "text":
          if (
              !/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/.test(value)
          ) {
              return "Not a valid email address";
          }
          break;
      case "password":
          if (value.length < 6) {
              return "Password is too short";
          }
          break;
      case "passwordConfirm":
          if (formData.password !== value) {
              return "Passwords don't match";
          }
          break;
  }
  return "";
};*/

const handleCheckboxChange = (event) => {
  const { name, checked } = event.target;

  

  setFormData((prevFormData) => {
    if (checked) {
      return {
        ...prevFormData,
        followers: [...prevFormData.followers, name],
      };
    } else {
      return {
        ...prevFormData,
        followers: prevFormData.followers.filter(follower => follower !== name),
      };
    }
  });
};



const handleSubmit = async (event) => {
  event.preventDefault();

  const payload = new FormData();

  for (const key in formData) {
      payload.append(key, formData[key]);
  }

  if (imageFile) {
      payload.append("image", imageFile);
  }

  if (groupId) {
    payload.append("group_id", groupId);
  }

  const options = {
      method: "POST",
      credentials: 'include',
      body: payload,
  };

  try {
      const response = await fetch(
          `http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/post/create`,
          options
      );
      if (response.ok) {
          router.push("/");

          const data = await response.json();
          window.location.reload();
      } else {
          const statusMsg = await response.text();
          console.log(statusMsg);
      }
  } catch (error) {
      console.error(error);
  }
};

  return (
    <>
      <form
        className="form"
        id="signupForm"
        onSubmit={handleSubmit}
        style={{ borderRadius: "10px" }}
      >
        <h2>{groupId ? "Create New Group Post" : "Create New Post"}</h2>
        <label htmlFor="title">Title</label>
        <input
          id="title"
          name="title"
          type="text"
          value={formData.title}
          onChange={handleChange} required
        />

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

        { !groupId && (
          <>
            <label>Privacy</label>
            <select 
              id="Privacy" 
              name="privacy" 
              value={formData.privacy} onChange={handleChange} required>
                <option value="">Select Privacy</option>
                <option value="public">Public</option>
                <option value="private">Private</option>
            </select>
            {formData.privacy === "private" && (
              <div>
                <label>Followers allowed</label>
                {followersArray.length > 0 && followersArray[0] != "" ? (
                  <>
                    {followersArray.map(follower => (
                      <div key={follower}>
                        <input 
                          type="checkbox" 
                          id={follower} 
                          name={follower} 
                          value={follower} 
                          checked={formData.followers.includes(follower)} 
                          onChange={handleCheckboxChange} 
                        />
                        <label htmlFor={follower}>{follower}</label>
                      </div>
                    ))}
                  </>
                ) : (
                  <p>No followers</p>
                )}
              </div>
            )}
          </>
        )}
        <button className="button-27" type="submit">Submit</button>
      </form>
    </>
  );
};


