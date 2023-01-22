import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

import Button from "@mui/material/Button";

import Container from "@mui/material/Container";

import Box from "@mui/material/Box";

import { CoursesInterface } from "../models/ICourse";

import { QualificationsInterface } from "../models/IQualification";

import { DataGrid, GridColDef } from "@mui/x-data-grid";

import SendIcon from '@mui/icons-material/Send';

import FolderIcon from '@mui/icons-material/Folder';
import { Grid } from "@mui/material";

function Users() {

 const [users, setUsers] = React.useState<CoursesInterface[]>([]);


 const getUsers = async () => {

   const apiUrl = "http://localhost:8080/courses";

   const requestOptions = {

     method: "GET",

     headers: { "Content-Type": "application/json" },

   };


   fetch(apiUrl, requestOptions)

     .then((response) => response.json())

     .then((res) => {

       console.log(res.data);

       if (res.data) {

         setUsers(res.data);

       }

     });

 };


 const columns: GridColDef[] = [

   { field: "Course_ID", headerName: "รหัสหลักสูตร", width:100 },

   { field: "Course_Name", headerName: "ชื่อหลักสูตร", width: 100},

   { field: "Datetime", headerName: "วันที่เพิ่ม", width: 150 },

   { field: "Qualification_ID", headerName: "วุฒิ", width: 100 },

   { field: "Major_ID", headerName: "สาขา", width: 100 },


 ];


 useEffect(() => {

   getUsers();

 }, []);

 

 return (

   <div>

     <Container maxWidth="md">

       <Box

         display="flex"

         sx={{

           marginTop: 2,

         }}

       >

         <Box flexGrow={1}>

           <Typography

             component="h2"

             variant="h6"

             color="primary"

             gutterBottom

           >

            
            
            <Grid item xs={6} color="#0895F7" 
          sx={{  fontFamily : "LilyUPC" ,
           fontWeight : 'bold' ,fontSize:27}}>
            <FolderIcon sx={{  fontFamily : "LilyUPC"  ,fontSize:55, mb:-2}} />
            ข้อมูลหลักสูตร
            </Grid>

           </Typography>

         </Box>

         <Box>

           <Button component={RouterLink} to="/create" sx={{  fontFamily : "LilyUPC"  ,fontSize:25}}
            variant="contained" color="success">

            <SendIcon/>  สร้างข้อมูลหลักสูตร

            

           </Button>

         </Box>

       </Box>

       <div style={{ height: 400, width: "100%", marginTop: '20px'}}>

         <DataGrid

           rows={users}

           getRowId={(row) => row.Course_ID}

           columns={columns}

           pageSize={5}

           rowsPerPageOptions={[5]}

         />

       </div>

     </Container>

   </div>

 );

}


export default Users;