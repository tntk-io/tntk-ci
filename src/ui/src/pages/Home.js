import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button } from "@material-ui/core";
import { useDispatch, useSelector } from "react-redux";
import { LinearProgress } from "@material-ui/core";
import {apiConfig } from "../helpers/api";
import Layout from "../layouts/index";
import { openModal } from "../redux/actions/modal";
import { saveAs } from 'file-saver';

const Home = () => {
  const dispatch = useDispatch();
  const { user } = useSelector((state) => state.auth);
  const [files, setFiles] = useState(null);

  useEffect(() => {
    const requestOptions = {
      headers: {
        'Token': user.token
      }
    }

    axios.get(`${apiConfig.getFiles}`, requestOptions)
      .then((res) => {
        setFiles(res.data);
      })
  }, [user.token])

  const reload = () => {
    setTimeout(() => {
      const requestOptions = {
        headers: {
          'Token': user.token
        }
      }

      axios.get(`${apiConfig.getFiles}`, requestOptions)
        .then((res) => {
          setFiles(res.data);
        })
    }, 2000)
  };

  const addPdf = () => {
    dispatch(
      openModal({
        component: "AddPDF",
        props: {
          user: user,
          reload: reload
        }
      })
    );
  };

  const downloadPdf = (file) => {
    const requestOptions = {
      headers: {
        'Token': user.token
      },
      responseType: "blob"
    }
    axios.get(`${apiConfig.getFile}/${file}`, requestOptions)
      .then(response => {
        saveAs(response.data, file)
      })
  };

  const deletePdf = (file) => {
    const requestOptions = {
        headers: {
            'Token' : user.token
        },
    }
    axios.delete(`${apiConfig.deleteFile}/${file}`, requestOptions)
  }

  return (
    <Layout>
      <h1>Welcome, {user.username}!</h1>
      <h2 style={{ marginTop: 50 }}>Generate PDF</h2>
      <Button
        color="primary"
        variant="contained"
        onClick={addPdf}
        style={{
          width: 142,
          borderRadius: 12,
          marginRight: 180,
          marginBottom: 90,
        }}
      >
        Request new PDF
      </Button>
      <h3>Generated PDF files</h3>

      {files ? files.map((file, index) => {
        return (
          <ul>
            <li key={index}>
              {file}
              <Button
                style={{ marginLeft: 20 }}
                variant="contained"
                color="secondary"
                onClick={() => downloadPdf(file)}>
                Download
              </Button>
              <Button
                style={{ marginLeft: 20 }}
                variant="contained"
                color="secondary"
                onClick={() => deletePdf(file)}>
                Delete
              </Button>
            </li>
          </ul>
        )
      }) : <><h2>Not PDF</h2></>}

    </Layout>
  );
};

export default Home;
