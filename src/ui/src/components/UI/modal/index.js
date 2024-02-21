import React from "react";
import CloseIcon from "@material-ui/icons/Close";
import { makeStyles } from "@material-ui/core/styles";
import Modal from "@material-ui/core/Modal";
import { useSelector, useDispatch } from "react-redux";
import { closeModal } from "../../../redux/actions/modal";
import AddPDF from "../../add-pdf";

const useStyles = makeStyles((theme) => ({
  modal: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },

  backDrop: {
    filter: "blur(3px)",
    "-webkit-filter": "blur(2px)",
  },

  paper: {
    background: theme.palette.common.white,
    boxShadow: "12px 8px 32px rgba(46, 71, 97, 0.3)",
    borderRadius: theme.spacing(2),
    padding: 50,
    position: "relative",
  },

  close: {
    position: "absolute",
    top: 10,
    right: 10,
    zIndex: 2,
    cursor: "pointer",
  },
}));

const modalChildrens = (props) => ({
  AddPDF: <AddPDF {...props} />,
});

export default function CenterModal() {
  const open = useSelector((state) => state.modal.modal);
  const props = useSelector((state) => state.modal.modalProps);
  const component = useSelector((state) => state.modal.component);
  const dispatch = useDispatch();
  const classes = useStyles();

  const handleClose = () => {
    dispatch(closeModal());
  };

  const close = () => {
    dispatch(closeModal());
  };

  return (
    <Modal
      className={classes.modal}
      open={open}
      onClose={handleClose}
      BackdropProps={{
        classes: {
          root: classes.backDrop,
        },
      }}
    >
      <>
        <div className={classes.paper}>
          <div>{modalChildrens(props)[component]}</div>
          <CloseIcon className={classes.close} onClick={close} />
        </div>
      </>
    </Modal>
  );
}
