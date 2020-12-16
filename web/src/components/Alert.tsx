import React from "react"
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from "@material-ui/core"

export function Alert({
  action,
  actionText,
  cancel,
  cancelText,
  message,
  open,
  title,
}: {
  action: () => void
  actionText: string
  cancel: () => void
  cancelText: string
  message: string
  open: boolean
  title: string
}): JSX.Element {
  return (
    <Dialog fullWidth maxWidth="xl" onClose={cancel} open={open}>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>
        <DialogContentText>{message}</DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button color="primary" onClick={cancel}>
          {cancelText}
        </Button>
        <Button autoFocus color="primary" onClick={action} variant="contained">
          {actionText}
        </Button>
      </DialogActions>
    </Dialog>
  )
}
