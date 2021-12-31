const token = JSON.parse(localStorage.getItem('token'));

const axiosConfig = {
  headers: {
    ContentType: 'application/json',
    Accept: 'application/json',
    Authorization: token ? `Bearer ${token}` : '',
  },
};
export default axiosConfig;
