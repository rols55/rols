import React, { useState } from "react";
import EmojiPicker, { Theme } from "emoji-picker-react";
import styles from "./page.module.css";

const EmojiPickerButton = ({handleEmojiPicker, className}) => {
    const [showPicker, setShowPicker] = useState(false);
    
    const onEmojiClick = (emojiObject) => {
        handleEmojiPicker(emojiObject.emoji);
        setShowPicker(false);
    };

    return (
        <div
            className={className}
            onClick={(e) => {
                e.preventDefault();
                setShowPicker(!showPicker)
            }}
        >
            <svg fill="none" height="40" viewBox="0 0 20 20" width="40" xmlns="http://www.w3.org/2000/svg">
                <g fill="#ffffff">
                    <path d="m6.49435 8.0754c.04511-.29574.33839-.5754.75364-.5754s.70853.27966.75364.5754c.04164.27298.29669.46052.56968.41888.27298-.04164.46052-.29669.41888-.56968-.12731-.83464-.88576-1.4246-1.7422-1.4246s-1.61489.58996-1.7422 1.4246c-.04164.27299.1459.52804.41888.56968.27299.04164.52804-.1459.56968-.41888z"/>
                    <path d="m12.748 7.5c-.4153 0-.7085.27966-.7536.5754-.0417.27298-.2967.46052-.5697.41888s-.4606-.29669-.4189-.56968c.1273-.83464.8858-1.4246 1.7422-1.4246s1.6149.58996 1.7422 1.4246c.0416.27299-.1459.52804-.4189.56968s-.528-.1459-.5697-.41888c-.0451-.29574-.3384-.5754-.7536-.5754z"/>
                    <path d="m5.49536 10c-.14106 0-.27556.0596-.37034.1641-.09477.1045-.141.2441-.12729.3845.23864 2.4432 2.15652 4.4514 4.99763 4.4514 2.84104 0 4.75894-2.0082 4.99764-4.4514.0137-.1404-.0325-.28-.1273-.3845s-.2293-.1641-.3703-.1641zm4.5 4c-2.08169 0-3.51612-1.3028-3.9122-3h7.82444c-.3962 1.6972-1.8306 3-3.91224 3z"/>
                    <path d="m10 2c-4.41828 0-8 3.58172-8 8 0 4.4183 3.58172 8 8 8 4.4183 0 8-3.5817 8-8 0-4.41828-3.5817-8-8-8zm-7 8c0-3.86599 3.13401-7 7-7 3.866 0 7 3.13401 7 7 0 3.866-3.134 7-7 7-3.86599 0-7-3.134-7-7z"/>
                </g>
            </svg>
            {showPicker && (
            <div className={`${styles.floatEmoji} ${showPicker ? styles.show : ''}`}>
                <EmojiPicker
                    lazyLoadEmojis={true}
                    searchDisabled={true}
                    skinTonesDisabled={true}
                    theme={Theme.DARK}
                    onEmojiClick={onEmojiClick}
                />
            </div>
            )}
        </div>
    );
};

export default EmojiPickerButton;