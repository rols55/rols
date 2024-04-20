"use client";
import { Dialog, Transition } from "@headlessui/react";
import { Fragment } from "react";
import "../app/globals.css";
import Edit from "./profile/edit";

export default function MyModal({ isOpen, setIsOpen, profile }) {
  return (
    <>
      <Transition appear show={isOpen} as={Fragment}>
        <Dialog as="div" onClose={() => setIsOpen(false)} className="dialog">
          <Dialog.Panel className="dialog-panel">
            <Edit className="edit" profile={profile} setIsOpen={setIsOpen}/>
          </Dialog.Panel>
        </Dialog>
      </Transition>
    </>
  );
}
