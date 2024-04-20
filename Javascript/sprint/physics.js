const getAcceleration = ({ f, m, Δv, Δt, d, t }) =>
  f && m ? f / m : Δv && Δt ? Δv / Δt : d && t ? (2 * d) / (t * t) : "impossible";