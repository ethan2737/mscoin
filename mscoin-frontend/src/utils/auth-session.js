export function getStoredMember({
  storage = localStorage,
  memberKey = 'MEMBER'
} = {}) {
  const rawMember = storage.getItem(memberKey)

  if (!rawMember) {
    return null
  }

  try {
    return JSON.parse(rawMember)
  } catch (error) {
    storage.removeItem(memberKey)
    return null
  }
}

export function hasAuthenticatedSession({
  storage = localStorage,
  tokenKey = 'TOKEN',
  memberKey = 'MEMBER',
  store
} = {}) {
  const token = storage.getItem(tokenKey)

  if (!token) {
    return false
  }

  if (store?.getters?.isLogin) {
    return true
  }

  return !!getStoredMember({ storage, memberKey })
}

export async function establishAuthenticatedSession({
  loginData,
  axiosInstance,
  memberInfoUrl,
  storage = localStorage,
  tokenKey = 'TOKEN',
  memberKey = 'MEMBER',
  store
}) {
  const token = loginData?.token

  if (!token) {
    throw new Error('missing auth token')
  }

  storage.setItem(tokenKey, token)

  let member = loginData

  if (axiosInstance && memberInfoUrl) {
    try {
      const response = await axiosInstance.post(memberInfoUrl, {}, {
        headers: { 'x-auth-token': token }
      })

      if (response?.data?.code === 0 && response.data.data) {
        member = {
          ...loginData,
          ...response.data.data,
          token
        }
      }
    } catch (error) {
      // Some environments do not expose a follow-up member-info endpoint.
      // In that case keep the successful login payload as the session source.
    }
  }

  if (!member) {
    throw new Error('missing member payload')
  }

  storage.setItem(memberKey, JSON.stringify(member))
  store?.commit?.('setMember', member)

  return member
}

export function clearAuthenticatedSession({
  storage = localStorage,
  tokenKey = 'TOKEN',
  memberKey = 'MEMBER',
  store
} = {}) {
  storage.removeItem(tokenKey)
  storage.removeItem(memberKey)
  store?.commit?.('removeMember')
}
