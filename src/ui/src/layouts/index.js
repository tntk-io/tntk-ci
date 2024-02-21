import logOutImg from "../assets/img/log-out.svg";
import { useDispatch } from "react-redux";
import { userActions } from "../redux/actions/auth";
import { makeStyles } from "@material-ui/core/styles";
import Box from "@material-ui/core/Box";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";

const useStyles = makeStyles((theme) => ({
  "@global": {
    body: {
      background: theme.palette.common.bg,
      overflowY: "auto",
    },
  },

  container: {
    position: "relative",
    padding: theme.spacing(4),
    minHeight: "100vh",
  },

  handlers: {
    position: "absolute",
    top: 45,
    right: theme.spacing(4),
  },

  handlersRow: {
    display: "flex",
    alignItems: "center",
    gap: 15,
    position: "relative",
  },

  aside: {
    position: "sticky",
    top: theme.spacing(4),
    height: "100%",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    background: theme.palette.common.white,
    minHeight: 500,
    borderRadius: theme.spacing(2),
    padding: `${theme.spacing(5)}px ${theme.spacing(2)}px`,
    width: "8rem",
  },

  logo: {
    marginBottom: theme.spacing(5),
    width: "5em",
  },
  button: {
    position: "relative",
    height: 20,
    width: 25,
    // backgroundImage: `url(${bell})`,
    backgroundRepeat: "no-repeat",
    backgroundPosition: "center",
    backgroundColor: "transparent",
    border: "none",
    opacity: 0.8,
    cursor: "pointer",
  },
  notificationCircle: {
    position: "absolute",
    top: -7,
    right: -2,
    width: 8,
    height: 8,
    backgroundColor: "#EB5757",
    borderRadius: "50%",
  },
  notification: {
    maxWidth: 300,
    marginBottom: 15,
    backgroundColor: "#E7EDF3",
    padding: 10,
    borderRadius: 10,
    textAlign: "center",
  },
  notificationRow: {
    display: "flex",
    alignItems: "flex",
    justifyContent: "center",
    marginTop: 10,
    gap: 10,
    flexWrap: "wrap",
    textAlign: "center",
  },
  notificationTitle: {
    display: "flex",
    gap: 10,
  },
  notificationTextBold: {
    fontWeight: "bold",
  },
  popover: {
    "& .MuiPopover-paper": {
      marginTop: 10,
      backgroundColor: "#fff",
      padding: "16px 16px 0 16px",
      borderRadius: 12,
    },
  },
}));

function MainLayout({ children }) {
  const dispatch = useDispatch();
  const classes = useStyles();
 

  return (
    <Container maxWidth="xl">
      <Box className={classes.container} display="flex">
        <div className={classes.handlers}>
          <div className={classes.handlersRow}>             
            <img
              src={logOutImg}
              alt=""
              title="logout"
              className="pointer"
              onClick={() => {
                dispatch(userActions.logout());
              }}
            />             
          </div>
        </div>
        {/*<Box>*/}
        {/*  <Paper component="aside" className={classes.aside}>*/}
        {/*  </Paper>*/}
        {/*</Box>*/}

        <Box style={{ flexGrow: 1, marginLeft: 50 }}>
          <main>{children}</main>
        </Box>
      </Box>
    </Container>
  );
}

export default MainLayout;
