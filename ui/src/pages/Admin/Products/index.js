import React, { useState } from "react";
import axios from 'axios';

function Products() {
  const [selectedImage, setSelectedImage] = useState([]);
  const [categoryId, setCategoryId] = useState(null);
  const [productName, setProductName] = useState("");
  const [description, setDescription] = useState("");
  const [price, setPrice] = useState(null);
  const [quantity, setQuantity] = useState(null);
  //const [imagePath, setImagePath] = useState("");
  console.log("images: ", selectedImage)
  const onImageChange = (event) => {
    const files = event.target.files;
    const newImagesArray = Array.from(files);
  
    // Append new images to the existing array
    setSelectedImage((prevImages) => [...prevImages, ...newImagesArray]);
  };

  const onFormSubmit = async () => {
    const formData = new FormData();
    selectedImage.forEach((image, index) => {
      formData.append(`images`, image);
    });
    formData.append('category_id', categoryId);
    formData.append('product_name', productName);
    formData.append('description', description);
    formData.append('price', price);
    formData.append('quantity', quantity);
    // Convert FormData to a regular object for logging
    const formDataObject = {};
    formData.forEach((value, key) => {
      formDataObject[key] = value;
    });

    console.log(formDataObject);

    //request to insert products
    const getTokenFromLocalStorage = () => {
      return localStorage.getItem('token');
    };
    const result = await axios.post('http://localhost:8080/user/add-product', formData, {
      headers: {
        'Authorization': `Bearer ${getTokenFromLocalStorage()}`,
        'Content-Type': 'multipart/form-data'
      }
    });
    console.log(result.data);
  };

  return (
    <div>
      <div>
      {selectedImage.length > 0 && (
        <div>
          Images: {selectedImage.map((image, index) => (
            <div key={index}>{image.name}</div>
          ))}
        </div>
      )}
      </div>
      <br></br>
      <input type="file" onChange={onImageChange} />
      <input type="number" placeholder="Category ID" onChange={e => setCategoryId(e.target.value)} />
      <input type="text" placeholder="Product Name" onChange={e => setProductName(e.target.value)} />
      <input type="text" placeholder="Description" onChange={e => setDescription(e.target.value)} />
      <input type="number" placeholder="Price" onChange={e => setPrice(e.target.value)} />
      <input type="number" placeholder="Quantity" onChange={e => setQuantity(e.target.value)} />
      {/* <input type="text" placeholder="Image Path" onChange={e => setImagePath(e.target.value)} /> */}
      <button onClick={onFormSubmit}>Submit</button>
    </div>
  );
}
export default Products;
