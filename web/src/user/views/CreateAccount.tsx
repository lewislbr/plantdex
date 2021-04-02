import React, {ChangeEvent, useContext, useEffect, useState} from "react"
import {useHistory} from "react-router-dom"
import {Button, IconButton, InputAdornment, TextField, Typography} from "@material-ui/core"
import {Visibility, VisibilityOff} from "@material-ui/icons"
import {Error, Loading} from "../../shared/components"
import {logIn, signUp} from "../services/user"
import {CREATE_ACCOUNT} from "../constants/copy"
import {EMAIL_PATTERN, NAME_PATTERN, PASSWORD_PATTERN} from "../constants/user"
import {CANCEL} from "../../shared/constants/copy"
import {HTTPStatus} from "../../shared/constants/http"
import {AuthContext} from "../contexts/auth"

export function CreateAccount(): JSX.Element {
  const [errors, setErrors] = useState({
    name: false,
    email: false,
    password: false,
    http: "",
  })
  const [buttonDisabled, setButtonDisabled] = useState(true)
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const [name, setName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [showPassword, setShowPassword] = useState(false)
  const {setAuthenticatedUser} = useContext(AuthContext)
  const history = useHistory()
  const missingFields = !name || !email || !password
  const activeErrors = Object.values(errors).includes(true)

  useEffect(() => {
    if (missingFields || activeErrors) {
      setButtonDisabled(true)
    } else {
      setButtonDisabled(false)
    }
  }, [missingFields, activeErrors])

  function updateName(event: ChangeEvent<HTMLInputElement>): void {
    if (!NAME_PATTERN.test(event.target.value)) {
      setErrors((errors) => ({...errors, name: true}))
    } else {
      setErrors((errors) => ({...errors, name: false}))
    }

    setName(event.target.value)
  }

  function updateEmail(event: ChangeEvent<HTMLInputElement>): void {
    if (!EMAIL_PATTERN.test(event.target.value)) {
      setErrors((errors) => ({...errors, email: true}))
    } else {
      setErrors((errors) => ({...errors, email: false}))
    }

    setEmail(event.target.value)
  }

  function updatePassword(event: ChangeEvent<HTMLTextAreaElement>): void {
    if (!PASSWORD_PATTERN.test(event.target.value)) {
      setErrors((errors) => ({...errors, password: true}))
    } else {
      setErrors((errors) => ({...errors, password: false}))
    }

    setPassword(event.target.value)
  }

  function toggleShowPassword(): void {
    setShowPassword(!showPassword)
  }

  async function handleSignUp(): Promise<void> {
    setStatus(HTTPStatus.LOADING)

    try {
      await signUp({
        name,
        email,
        password,
      })
      await logIn({
        email,
        password,
      })

      setAuthenticatedUser(true)
      setStatus(HTTPStatus.SUCCESS)

      history.push("/plants")
    } catch (error) {
      setErrors((errors) => ({...errors, http: String(error)}))
      setStatus(HTTPStatus.ERROR)

      console.error(error)
    }
  }

  function cancel(): void {
    history.push("/")
  }

  return (
    <>
      <Typography variant="h1">{CREATE_ACCOUNT}</Typography>
      {status === HTTPStatus.LOADING ? (
        <Loading />
      ) : (
        <section style={{marginTop: "30px"}}>
          <TextField
            error={errors.name}
            fullWidth
            label="Name"
            onChange={updateName}
            required
            value={name}
            variant="outlined"
          />
          <TextField
            error={errors.email}
            fullWidth
            label="Email"
            onChange={updateEmail}
            required
            type="email"
            value={email}
            variant="outlined"
          />
          <TextField
            error={errors.password}
            fullWidth
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton edge="end" onClick={toggleShowPassword}>
                    {showPassword ? <Visibility /> : <VisibilityOff />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
            label="Password (min. 8 characters)"
            onChange={updatePassword}
            required
            type={showPassword ? "text" : "password"}
            value={password}
            variant="outlined"
          />
          <Button
            color="primary"
            disabled={buttonDisabled}
            fullWidth
            onClick={handleSignUp}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {CREATE_ACCOUNT}
          </Button>
          <Button
            color="secondary"
            fullWidth
            onClick={cancel}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {CANCEL}
          </Button>
        </section>
      )}
      {status === HTTPStatus.ERROR && <Error message={errors.http} title={"Error"} />}
    </>
  )
}
