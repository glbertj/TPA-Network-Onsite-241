import type { Dispatch, SetStateAction } from "react";
import { useEffect, useState } from "react";

interface SuccessModalProps {
  success: string;
  setSuccess: Dispatch<SetStateAction<string>>;
}

export const SuccessModal = ({ success, setSuccess }: SuccessModalProps) => {
  const [showModal, setShowModal] = useState(false);

  useEffect(() => {
    let timeout: number;

    if (success) {
      setShowModal(true);
      timeout = window.setTimeout(() => {
        setShowModal(false);
        setSuccess("");
      }, 3000);
    }

    return () => {
      clearTimeout(timeout);
    };
  }, [success, setSuccess]);

  const closeModal = () => {
    setShowModal(false);
  };

  return (
    <div className="error-modal">
      {showModal && (
        <div className="modal">
          <div className="modal-content">
            <span className="close" onClick={closeModal}>
              &times;
            </span>
            <p className={"success"}>{success}</p>
          </div>
        </div>
      )}
    </div>
  );
};
