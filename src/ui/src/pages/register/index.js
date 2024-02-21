import { useState, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { NavLink } from 'react-router-dom';
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import Box from "@material-ui/core/Box";
import { userActions } from "../../redux/actions/auth";
import CircularProgress from "@material-ui/core/CircularProgress";
import { makeStyles } from "@material-ui/core/styles";
import bg from "../../assets/img/bg.png";

const useStyles = makeStyles((theme) => ({
    container: {
        background: `url(${bg}) no-repeat`,
        backgroundSize: "cover",
        display: "flex",
        padding: "2.5vh",
        position: "relative",
    },
    title: {
        fontSize: 44,
        fontWeight: 700,
        maxWidth: 300,
        marginBottom: 20,
        lineHeight: "52px",
    },
    form: {
        background: "#fff",
        width: 480,
        height: "95vh",
        minHeight: 700,
        borderRadius: 16,
        marginLeft: "auto",
        marginRight: "2.5%",
        padding: "70px 50px 40px",
        display: "flex",
        flexDirection: "column",
    },
    errorMessage: {
        fontSize: 12,
        fontWeight: 500,
        color: "#eb5757",
        marginTop: 15,
    },

    textFieldRoot: {
        backgroundColor: theme.palette.common.inputbg,
        borderRadius: 12,
        "&.Mui-error .MuiOutlinedInput-notchedOutline": {
            border: "1px solid #f44336",
        },
    },

    btn: {
        marginTop: "auto",
        width: 200,
        alignSelf: "flex-end",
        boxShadow: "none",
        color: "#fff"
    },

    focused: {
        backgroundColor: theme.palette.common.inputbg,
    },

    label: {
        color: "#000",
    },

    notchedOutline: {
        border: "none",
    },

    error: {
        color: "#f44336",
    },
}));

function Register() {
    const [values, setValues] = useState({
        password: "",
        username: "",
    });
    const [isValidData, setIsValidData] = useState(true);
    const dispatch = useDispatch();
    const loginError = useSelector((state) => state.auth.error);
    const loading = useSelector((state) => state.auth.loading);
    const isPasswordValid = values.password.length > 5;
    const styles = useStyles();

    useEffect(() => {
        if (loginError) {
            setIsValidData(false);
        }
    }, [loginError]);


    const onSubmit = (e) => {
        e.preventDefault();
        const { username, password } = values;

        if (isPasswordValid) {
            dispatch(userActions.register(username, password));
            return;
        }
        setIsValidData(false);
    };

    const hangleChange = (e) => {
        const { value, name } = e.target;
        setValues((prev) => ({ ...prev, [name]: value }));
    };

    const onFocus = () => {
        setIsValidData(true);
    };

    return (
        <Box className={styles.container}>
            <form className={styles.form} onSubmit={onSubmit} noValidate>
                <h1 className={styles.title}>Welcome back!</h1>
                <p className="caption" style={{ marginBottom: 50 }}>
                    Please enter your username and password.
                </p>
                <Box>
                    <TextField
                        label="Username"
                        name="username"
                        fullWidth
                        variant="outlined"
                        style={{ marginBottom: 20 }}
                        value={values.username}
                        onChange={hangleChange}
                        onFocus={onFocus}
                        error={!isValidData}
                        InputProps={{
                            classes: {
                                root: styles.textFieldRoot,
                                focused: styles.focused,
                                notchedOutline: styles.notchedOutline,
                                error: styles.error,
                            },
                        }}
                        InputLabelProps={{
                            classes: { root: styles.label },
                        }}
                    />

                    <TextField
                        label="Password"
                        name="password"
                        type="password"
                        fullWidth
                        variant="outlined"
                        style={{ marginBottom: 20 }}
                        value={values.password}
                        onChange={hangleChange}
                        onFocus={onFocus}
                        error={!isValidData}
                        InputProps={{
                            classes: {
                                root: styles.textFieldRoot,
                                focused: styles.focused,
                                notchedOutline: styles.notchedOutline,
                                error: styles.error,
                            },
                        }}
                        InputLabelProps={{
                            classes: { root: styles.label },
                        }}
                    />

                    {!isValidData ? (
                        <div className={styles.errorMessage}>
                            {loginError}
                        </div>
                    ) : null}
                </Box>
                <div>
                    Have account? <NavLink to="/login">Enter</NavLink>
                </div>
                <Button
                    type="submit"
                    color={values.username && values.password ? "primary" : "secondary"}
                    size="large"
                    variant="contained"
                    className={styles.btn}
                >
                    {loading ? <CircularProgress size={16} /> : "Sign Up"}
                </Button>
            </form>
        </Box>
    );
}

export default Register;
