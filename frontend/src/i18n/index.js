import en from '../locales/en.json';

const locales = { en };
let currentLocale = 'en';

/**
 * Set the active locale.
 * @param {string} locale
 */
export function setLocale(locale) {
  if (!locales[locale]) {
    console.warn(`[i18n] Unknown locale "${locale}", falling back to "en".`);
    currentLocale = 'en';
    return;
  }
  currentLocale = locale;
}

/**
 * Get the active locale code.
 * @returns {string}
 */
export function getLocale() {
  return currentLocale;
}

/**
 * Translate a key using the current locale.
 * Returns the key itself if no translation is found.
 * @param {string} key
 * @returns {string}
 */
export function t(key) {
  const dict = locales[currentLocale] || locales.en;
  return dict[key] !== undefined ? dict[key] : key;
}

/**
 * Register additional locale data at runtime.
 * @param {string} locale
 * @param {Record<string, string>} messages
 */
export function registerLocale(locale, messages) {
  locales[locale] = { ...(locales[locale] || {}), ...messages };
}
