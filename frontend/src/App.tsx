import React from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import UserInfo from "./components/UserInfo";

import UserCreate from "./components/UserCreate";

import UserList from "./components/UserList";

export default function App() {

return (

  <Router>

   <div>

   <Navbar />

   <Routes>

       <Route path="/" element={<UserList />} />

       <Route path="/create" element={<UserCreate />} />

       <Route path="/info/:course_id" element={<UserInfo />} />

   </Routes>

   </div>

  </Router>

);

}
